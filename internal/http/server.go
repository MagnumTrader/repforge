package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MagnumTrader/repforge/internal/http/handlers"
	"github.com/MagnumTrader/repforge/internal/http/routes"
	"github.com/MagnumTrader/repforge/internal/http/ui"
	"github.com/MagnumTrader/repforge/internal/infrastructure/db"
	"github.com/MagnumTrader/repforge/internal/services"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func mainPage(ctx *gin.Context) {
	var template templ.Component = ui.MainPage()
	if !handlers.IsHtmxRequest(ctx) {
		// Fresh page load
		template = ui.Base(template)
	}
	template.Render(ctx.Request.Context(), ctx.Writer)
}

func healthyHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "healthy")
}

func GetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Static("/static", "./internal/http/static")

	db := db.NewDb()
	service := services.NewWorkoutService(db)
	routes.RegisterWorkoutRoutes(r, service)

	r.GET("/", mainPage)
	r.GET("/health", healthyHandler)
	r.GET("/time", func(ctx *gin.Context) {
		time := time.Now()
		s := time.Format("15:04")

		fmt.Fprintf(ctx.Writer, "%s", s)
	})

	return r
}
