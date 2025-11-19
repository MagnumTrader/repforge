package main

import (
	"fmt"
	"log"
	"strconv"

	shttp "net/http"

	"github.com/MagnumTrader/repforge/internal/config"
	"github.com/MagnumTrader/repforge/internal/hotreload"
	"github.com/MagnumTrader/repforge/internal/http"
)

func main() {

	r := http.GetRouter()

	c := hotreload.RegisterWatcher("./internal/http/static", "./internal/http/static/styles")
	r.GET("/hotreload", hotreload.HotreloadHandler(c))

	addr := "127.0.0.1:" + strconv.Itoa(config.Port)
	fmt.Printf("Listening on http://%s\n", addr)

	log.Fatal(shttp.ListenAndServe(addr, r))
}
