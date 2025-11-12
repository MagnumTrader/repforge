package hotreload

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func RegisterWatcher(paths ...string) <-chan int8 {

	if watcher != nil {
		panic("Can only register one time!")
	}

	watcher, _ = fsnotify.NewWatcher()

	for _, path := range paths {
		watcher.Add(path)
	}

	fmt.Printf("Watcher is watching: %v", watcher.WatchList())

	c := make(chan int8)
	go func() {
		for {
			msg := <-watcher.Events

			if msg.Op == fsnotify.Rename {
				fmt.Printf("File changed %s", msg.Name)
				c <- 1
				watcher.Add(msg.Name)
			}
		}
	}()
	return c
}
