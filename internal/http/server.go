package http

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/MagnumTrader/repforge/internal/domain"
	"github.com/MagnumTrader/repforge/internal/http/ui"
	"github.com/MagnumTrader/repforge/internal/infrastructure/db"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type app struct {
	db *db.Db
}

const (
	HXREQUEST    = "Hx-Request"
	HXCURRENTURL = "Hx-Current-Url"
)

func (app *app) mainPage(ctx *gin.Context) {
	var template templ.Component = ui.MainPage()
	if !isHtmxRequest(ctx) {
		slog.Info("request came in", "context", ctx.Request)
		template = ui.Base(template)
	}
	template.Render(ctx.Request.Context(), ctx.Writer)
}

func (app *app) workoutsListHandler(ctx *gin.Context) {
	workouts, err := app.db.GetAllWorkouts(0)
	if err != nil {
		slog.Error(err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}
	template := ui.WorkOutListPartial(workouts)

	if !isHtmxRequest(ctx) {
		template = ui.Base(template)
	}

	template.Render(ctx.Request.Context(), ctx.Writer)
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

	template := ui.WorkoutDetailsPartial(*workout)
	if !isHtmxRequest(ctx) {
		template = ui.Base(template)
	}
	template.Render(ctx.Request.Context(), ctx.Writer)
}
func (app *app) newWorkoutForm(ctx *gin.Context) {
	if isHtmxRequest(ctx) {
		// we should render the partial
		template := ui.WorkoutForm(domain.Workout{})
		template.Render(ctx.Request.Context(), ctx.Writer)
		return
	}
}

func (app *app) newWorkout(ctx *gin.Context) {

	workout := struct {
		Date     string `form:"date"  binding:"required"`
		Duration int    `form:"duration"`
		Type     string `form:"type" binding:"required"`
		Note     string `form:"note" `
	}{}

	if err := ctx.ShouldBind(&workout); err != nil {
		slog.Error("Failed to parse form", "Error:", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	slog.Info("created entity", "Name: ", workout.Date)
	slog.Info("", "email: ", workout.Note)

	// here we should probablt have a service to the handler?
	wo := domain.Workout{
		Date:     workout.Date,
		Type:     workout.Type,
		Duration: workout.Duration,
		Notes:    workout.Note,
	}

	err := app.db.SaveWorkout(&wo)

	if err != nil {
		slog.Error("Failed to insert workout into DB", "Error:", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	row := ui.WorkoutTableRow(wo)
	row.Render(ctx.Request.Context(), ctx.Writer)
}

func isHtmxRequest(ctx *gin.Context) bool {
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
	r.GET("/workouts/new", app.newWorkoutForm)
	r.POST("/workouts/new", app.newWorkout)
	r.GET("/time", func(ctx *gin.Context) {
		time := time.Now()
		s := time.Format("15:04")

		fmt.Fprintf(ctx.Writer, "%s", s)
	})

	return r
}
