package http

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/MagnumTrader/repforge/internal/domain"
	"github.com/MagnumTrader/repforge/internal/http/routes"
	"github.com/MagnumTrader/repforge/internal/http/ui"
	"github.com/MagnumTrader/repforge/internal/infrastructure/db"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type app struct {
	db *db.Db
}

func (app *app) mainPage(ctx *gin.Context) {
	var template templ.Component = ui.MainPage()
	if !isHtmxRequest(ctx) {
		slog.Info("request came in", "context", ctx.Request)
		template = ui.Base(template)
	}
	template.Render(ctx.Request.Context(), ctx.Writer)
}


func GetRouter() *gin.Engine {
	app := &app{
		db: db.NewDb(),
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Static("/static", "./internal/http/static")

	//TODO: Here we should register the workout handler instead

	routes.WorkoutRoutes()

	r.GET("/", app.mainPage)
	r.GET("/health", app.healthyHandler)
	r.GET("/workouts", app.workoutsListHandler)
	r.GET("/workouts/:id", app.workoutDetails)
	r.DELETE("/workouts/:id", app.deleteWorkout)
	r.GET("/workouts/new", app.newWorkoutForm)
	r.POST("/workouts/new", app.newWorkout)
	r.GET("/time", func(ctx *gin.Context) {
		time := time.Now()
		s := time.Format("15:04")

		fmt.Fprintf(ctx.Writer, "%s", s)
	})

	return r
}
