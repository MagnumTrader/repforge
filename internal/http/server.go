package http

import (
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/MagnumTrader/repforge/internal/http/ui"
	"github.com/MagnumTrader/repforge/internal/infrastructure/db"
	"github.com/gin-gonic/gin"
)

// So we got off topic
// But now we can get to work. what is the next thing?


type app struct {
	db *db.Db
}

func (app *app) mainPage(ctx *gin.Context) {
	ui.MainPage().Render(ctx.Request.Context(), ctx.Writer)
}

func (app *app) workoutsListHandler(ctx *gin.Context) {
	workouts, err := app.db.GetAllWorkouts(0)
	if err != nil {
		slog.Error(err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ui.WorkOutList(workouts).Render(ctx.Request.Context(), ctx.Writer)

}
func (app *app) healthyHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "this is my and this now changed")
}
func (app *app) workoutDetails(ctx *gin.Context) {
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

	workout, err := app.db.GetWorkout(parsedId)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ui.WorkoutDetailPage(*workout).Render(ctx.Request.Context(), ctx.Writer)
}

func GetRouter() *gin.Engine {
	app := &app{ 
		db: db.NewDb(),
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Static("/static", "./internal/http/static")

	r.GET("/", app.mainPage)
	r.GET("/health", app.healthyHandler)
	r.GET("/workouts", app.workoutsListHandler)
	r.GET("/workouts/:id", app.workoutDetails)

	return r
}
