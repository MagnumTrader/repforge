package http

import (
	"net/http"

	"github.com/MagnumTrader/repforge/internal/domain"
	"github.com/MagnumTrader/repforge/internal/ui"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ui.MainPage().Render(ctx.Request.Context(), ctx.Writer)
	})
	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "healthy")
	})
	r.GET("/workouts", func(ctx *gin.Context) {
		ui.WorkOutList(domain.Workouts).Render(ctx.Request.Context(), ctx.Writer)
	})
	return r
}
