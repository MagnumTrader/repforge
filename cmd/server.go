package main

import (
	"log"

	"github.com/MagnumTrader/repforge/internal/config"
	"github.com/MagnumTrader/repforge/internal/http"
)

func main()  {
	log.Fatal(http.RunServer(config.Port))
}

