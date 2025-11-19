package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/MagnumTrader/repforge/internal/http/ui"
	"github.com/MagnumTrader/repforge/internal/services"

	"github.com/gin-gonic/gin"
)

type workout struct {
	service *services.WorkoutService
}

func NewWorkoutHandler(service *services.WorkoutService) workout {
	return workout{
		service: service,
	}

}

func (h *workout) WorkoutsList(ctx *gin.Context) {
	workouts, err := h.service.GetAll()
	if err != nil {
		slog.Error(err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}
	template := ui.WorkOutListPartial(workouts)

	if !IsHtmxRequest(ctx) {
		template = ui.Base(template)
	}

	template.Render(ctx.Request.Context(), ctx.Writer)
}

func (h *workout) WorkoutDetails(ctx *gin.Context) {
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

	workout, err := h.service.GetWorkout(parsedId)
	if err != nil {
		slog.Info("failed to get workout2", "error", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	template := ui.WorkoutDetailsPartial(*workout)
	if !IsHtmxRequest(ctx) {
		template = ui.Base(template)
	}
	template.Render(ctx.Request.Context(), ctx.Writer)
}

func (h *workout) FormNewWorkout(ctx *gin.Context) {
	if IsHtmxRequest(ctx) {
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
	err = h.service.DeleteWorkout(parsedId)
	if err != nil {
		slog.Error("Error when deleting workout", "error", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
	ctx.Writer.Write([]byte(""))
}

// The handler should have its own request format, it has that here
// This is for the post handling of that request
func (h *workout) NewWorkout(ctx *gin.Context) {

	wo := struct {
		Date     string `form:"date"  binding:"required"`
		Duration int    `form:"duration"`
		Type     string `form:"type" binding:"required"`
		Note     string `form:"note" `
	}{}

	if err := ctx.ShouldBind(&wo); err != nil {
		slog.Error("Failed to parse form", "Error:", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	workout, err := h.service.CreateWorkout(wo.Date, wo.Type, wo.Note, wo.Duration)

	if err != nil {
		slog.Error("Failed to insert workout into DB", "Error:", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	row := ui.WorkoutTableRow(*workout)
	row.Render(ctx.Request.Context(), ctx.Writer)
}
