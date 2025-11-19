package hotreload

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

func HotreloadHandler(c <-chan string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "text/event-stream")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Connection", "keep-alive")
		ctx.Writer.Header().Set("Transfer-Encoding", "chunked")
		for {
			select {
			case file := <-c:
				fmt.Println("Hotreload triggered")
				if _, err := fmt.Fprintf(ctx.Writer, "data: file %s has changed\n\n", file); err != nil {
					fmt.Println(err)
					return
				}
				ctx.Writer.Flush()
			case <-ctx.Request.Context().Done():
				return
			}
		}
	}
}

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
