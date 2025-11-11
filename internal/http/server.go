package http

import (
	"fmt"
	"net/http"
	"strconv"

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
	r.GET("/workouts/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		fmt.Printf("id is: %s", id)
		if id == "" {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		parsedId, err:= strconv.Atoi(id)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		for _, wo := range domain.Workouts {
			if wo.Id == parsedId {
				ui.WorkoutDetailPage(wo).Render(ctx.Request.Context(), ctx.Writer)
				return
			}
		}
		
	})
	return r
}
