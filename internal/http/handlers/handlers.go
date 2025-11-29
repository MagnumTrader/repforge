package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

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

// Parse the id of the request, MAY be more general later if we have structure for system wide Id's
func parseId(ctx *gin.Context) (int, error) {
	idString := ctx.Param("id")
	if idString == "" {
		return 0, fmt.Errorf("%w: missing", ErrBadID)
	}

	id, err := strconv.Atoi(idString)

	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrBadID, idString)
	}
	return id, nil
}
