package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/MagnumTrader/repforge/internal/domain"
	"github.com/MagnumTrader/repforge/internal/http/ui"
	"github.com/gin-gonic/gin"
	"github.com/gohugoio/hugo/config/services"
)

type workout struct {
	service services.WorkoutService
}

func NewWorkoutHandler(service services.WorkoutService) workout {
	return workout{
		service: service,
	}

}

func (h *workout) WorkoutsListHandler(ctx *gin.Context) {
	workouts, err := h.service.GetAll()
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

func (h *workout) healthyHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "this is my and this now changed")
}

func (h *workout) workoutDetails(ctx *gin.Context) {
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

func (h *workout) WorkoutForm(ctx *gin.Context) {
	if isHtmxRequest(ctx) {
		// we should render the partial
		template := ui.WorkoutForm(nil)
		template.Render(ctx.Request.Context(), ctx.Writer)
		return
	}
}

func (h *workout) DeleteWorkout(ctx *gin.Context) {
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
	err = app.db.DeleteWorkout(parsedId)
	if err != nil {
		slog.Error("Error when deleting workout", "error", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
	ctx.Writer.Write([]byte(""))
}

// The handler should have its own request format, it has that here
func (h *workout) NewWorkout(ctx *gin.Context) {

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

