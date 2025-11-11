package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RunServer(port int) error  {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	addr := "127.0.0.1:" + strconv.Itoa(port)
	fmt.Printf("Listening on http://%s", addr)
	return http.ListenAndServe(addr, r)
}




