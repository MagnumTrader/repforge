package handlers

import "github.com/gin-gonic/gin"

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

func headerExist(ctx *gin.Context, header string) bool {
	headers := ctx.Request.Header
	_, ok := headers[header]
	return ok
}
