package main

import (
	"fmt"
	"log"
	"strconv"

	shttp "net/http"

	"github.com/MagnumTrader/repforge/internal/config"
	"github.com/MagnumTrader/repforge/internal/hotreload"
	"github.com/MagnumTrader/repforge/internal/http"
	"github.com/gin-gonic/gin"
)

func main()  {

	r := http.GetRouter()

	c := hotreload.RegisterWatcher("./internal/http/static", "./internal/http/static/styles")
	r.GET("/hotreload", func(ctx *gin.Context) {
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
	})

	addr := "127.0.0.1:" + strconv.Itoa(config.Port)
	fmt.Printf("Listening on http://%s\n", addr)
	
	log.Fatal(shttp.ListenAndServe(addr, r))
}

