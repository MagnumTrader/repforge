package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HXREQUEST    = "Hx-Request"
	HXCURRENTURL = "Hx-Current-Url"
)

func IsHtmxRequest(ctx *gin.Context) bool {
	value, ok := ctx.Request.Header[HXREQUEST]
	if ok {
		return value[0] == "true"
	}
	return false
}

// Sends back a response with a status and msg
// logs the error on the server for troubleshooting
func respondError(ctx *gin.Context, status int, msg string, err error) {
	slog.Error(msg, "error", err)
	ctx.String(status, msg)
}

// Sets the content type to text/html 
// and response to be 200
func setHtml200(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.Status(http.StatusOK)
}
