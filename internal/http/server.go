package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MagnumTrader/repforge/internal/config"
	"github.com/MagnumTrader/repforge/internal/domain"
	"github.com/MagnumTrader/repforge/internal/hotreload"
	"github.com/MagnumTrader/repforge/internal/http/ui"
	"github.com/gin-gonic/gin"
)

// So we got off topic
// But now we can get to work. what is the next thing?

// TODO:
// - [ ] Add sql folder structure
// - [ ] Add table for workouts AS IS
// - [ ] Add functions for creating a workout and store in db
// - [ ] Append a new row to the list of workouts
// - [ ] remove notes in list view, add Score ( your own score and what you thought)

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

	c := hotreload.RegisterWatcher("./internal/http/static", "./internal/http/static/styles")

	r.GET("/hotreload", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "text/event-stream")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Connection", "keep-alive")
		ctx.Writer.Header().Set("Transfer-Encoding", "chunked")

		for {
			select {
			case file := <-c:
				if _, err := fmt.Fprintf(ctx.Writer, "data: file %s has changed\n\n", file); err != nil {
					fmt.Println(err)
					return
				}
				ctx.Writer.Flush()
			case <-ctx.Request.Context().Done():
				return
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
		if id == "" {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		parsedId, err := strconv.Atoi(id)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		// this logic
		for _, wo := range domain.Workouts {
			if wo.Id == parsedId {
				ui.WorkoutDetailPage(wo).Render(ctx.Request.Context(), ctx.Writer)
				return
			}
		}
	})
	return r
}
