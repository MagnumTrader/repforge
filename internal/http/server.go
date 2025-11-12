package http

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/MagnumTrader/repforge/internal/config"
	"github.com/MagnumTrader/repforge/internal/domain"
	"github.com/MagnumTrader/repforge/internal/ui"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// NOTE: using this to remove messages used by hotreloading
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/version"},
	}))
	r.Use(gin.Recovery())

	r.Static("/static", "./internal/http/static")

	r.GET("/", func(ctx *gin.Context) {
		ui.MainPage().Render(ctx.Request.Context(), ctx.Writer)
	})
	r.GET("/sse", func(c *gin.Context) {
		// TODO: so, now we have the 
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")

		// Get the client's context

		// Keep connection alive and send events
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-c.Request.Context().Done():
				// Client disconnected
				fmt.Println("client disconnected from SSE")
				return
			case t := <-ticker.C:
				// Send a message every second
				msg := fmt.Sprintf("data: The time is %s\n\n", t.Format(time.RFC3339))
				_, err := io.WriteString(c.Writer, msg)
				if err != nil {
					return
				}
				c.Writer.Flush()
			}
		}
	})

	r.GET("/version", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, config.Version)
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

		parsedId, err := strconv.Atoi(id)

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
