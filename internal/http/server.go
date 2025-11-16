package http

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/MagnumTrader/repforge/internal/http/ui"
	"github.com/MagnumTrader/repforge/internal/infrastructure/db"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

// So we got off topic
// But now we can get to work. what is the next thing?
/*
TODO:
Lets move this shit to htmx.
The main things are to separate partials and site
- [x]  make the base generic with a main #content
- [x]  remake other pages to be partials
- [x]  lets do so we can use it as is but with this structure
*/

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
func (app *app) newWorkout(ctx *gin.Context) {
	if headerExist(ctx, HXREQUEST) {
		// we should render the partial
		slog.Info("this is a htmx request")
		ctx.Writer.Write([]byte(`<span hx-on:click="alert('hello world i said!')">'hello world'</span>`))
	} else {
		slog.Info("this is not")
	}
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
	r.GET("/workouts/new", app.newWorkout)
	r.GET("/time", func(ctx *gin.Context) {
		time := time.Now()
		s := time.Format("15:04")

		fmt.Fprintf(ctx.Writer, "%s", s)
	})

	return r
}
