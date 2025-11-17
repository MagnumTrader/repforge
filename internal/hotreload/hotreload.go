package hotreload

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func RegisterWatcher(paths ...string) <-chan string {

	if watcher != nil {
		panic("Can only register hotreload one time!")
	}

	watcher, _ = fsnotify.NewWatcher()

	for _, path := range paths {
		watcher.Add(path)
	}

	fmt.Print("Watcher is watching: ")
	for _, file := range watcher.WatchList() {
		fmt.Printf(`"%s", `, file)
	}
	fmt.Println()

	c := make(chan string, 500)
	go func() {
		for {
			msg := <-watcher.Events
			if msg.Op == fsnotify.Rename {
				c <- msg.Name
				watcher.Add(msg.Name)
			}
		}
	}()
	return c
}
