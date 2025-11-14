package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/MagnumTrader/repforge/internal/config"
	"github.com/MagnumTrader/repforge/internal/http/ui"
	"github.com/MagnumTrader/repforge/internal/infrastructure/db"
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

var woRepo = db.InMem{}

func GetRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Static("/static", "./internal/http/static")

	r.GET("/", func(ctx *gin.Context) {
		ui.MainPage().Render(ctx.Request.Context(), ctx.Writer)
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "healthy")
	})
	r.GET("/workouts", func(ctx *gin.Context) {

		workouts, err := woRepo.GetAllWorkouts(0)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ui.WorkOutList(workouts).Render(ctx.Request.Context(), ctx.Writer)
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

		workout, err := woRepo.GetWorkout(parsedId)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ui.WorkoutDetailPage(*workout).Render(ctx.Request.Context(), ctx.Writer)
	})

	return r
}
