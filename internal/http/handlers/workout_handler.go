package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/MagnumTrader/repforge/internal/domain"
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
		respondError(ctx, http.StatusInternalServerError, FailedToGetWorkouts, err)
		return
	}
	template := ui.WorkOutListPartial(workouts)

	setHtml200(ctx)

	if !IsHtmxRequest(ctx) {
		template = ui.Base(template)
	}
	template.Render(ctx.Request.Context(), ctx.Writer)
}

func (h *workout) WorkoutDetails(ctx *gin.Context) {
	workout, err := h.getWorkoutByCtxId(ctx)
	if err != nil {
		if errors.Is(err, ErrBadID) {
			respondError(ctx, http.StatusBadRequest, InvalidWorkoutId, err)
			return
		}
		respondError(ctx, http.StatusInternalServerError, FailedToGetWorkout, err)
		return
	}

	template := ui.WorkoutDetailsPartial(*workout)

	setHtml200(ctx)

	if !IsHtmxRequest(ctx) {
		template = ui.Base(template)
	}
	template.Render(ctx.Request.Context(), ctx.Writer)
}

func (h *workout) EditWorkoutForm(ctx *gin.Context) {
	wo, err := h.getWorkoutByCtxId(ctx)

	if err != nil {
		if errors.Is(err, ErrBadID) {
			respondError(ctx, http.StatusBadRequest, InvalidWorkoutId, err)
			return
		}
		respondError(ctx, http.StatusInternalServerError, FailedToGetWorkout, err)
		return
	}

	setHtml200(ctx)

	if IsHtmxRequest(ctx) {
		// We should render the partial
		template := ui.WorkoutForm(wo, ui.EditForm)
		template.Render(ctx.Request.Context(), ctx.Writer)
		return
	}

	// NOTE: maybe support standalone screen later
	ctx.Status(http.StatusNotFound)
}

func (h *workout) EditWorkout(ctx *gin.Context) {

	wo := struct {
		Date     string `form:"date"  binding:"required"`
		Duration int    `form:"duration"`
		Type     string `form:"type" binding:"required"`
		Note     string `form:"note" `
	}{}

	if err := ctx.ShouldBind(&wo); err != nil {
		respondError(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	id, err := parseId(ctx)
	if err != nil {
		respondError(ctx, http.StatusInternalServerError, InvalidWorkoutId, err)
	}

	// NOTE: Maybe use DTO later
	workout := &domain.Workout{
		Id:       id,
		Date:     wo.Date,
		Kind:     wo.Type,
		Duration: wo.Duration,
		Notes:    wo.Note,
	}

	err = h.service.EditWorkout(workout)

	if err != nil {
		respondError(ctx, http.StatusInternalServerError, "Failed to update workout", err)
		return
	}

	setHtml200(ctx)
	template := ui.WorkoutTableRow(*workout)
	template.Render(ctx.Request.Context(), ctx.Writer)
}

func (h *workout) NewWorkoutForm(ctx *gin.Context) {
	if IsHtmxRequest(ctx) {
		// we should render the partial
		template := ui.WorkoutForm(nil, ui.NewForm)
		setHtml200(ctx)
		template.Render(ctx.Request.Context(), ctx.Writer)
		return
	}

	// NOTE: maybe support standalone screen later
	ctx.Status(http.StatusNotFound)
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
		respondError(ctx, http.StatusBadRequest, InvalidInput, err)
		return
	}

	workout, err := h.service.CreateWorkout(wo.Date, wo.Type, wo.Note, wo.Duration)

	if err != nil {
		respondError(ctx, http.StatusInternalServerError, InvalidInput, err)
		slog.Error("Failed to insert workout into DB", "Error:", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	row := ui.WorkoutTableRow(*workout)
	setHtml200(ctx)
	row.Render(ctx.Request.Context(), ctx.Writer)
}

func (h *workout) DeleteWorkout(ctx *gin.Context) {

	id, err := parseId(ctx)
	if err != nil {
		respondError(ctx, http.StatusBadRequest, InvalidWorkoutId, err)
		return
	}

	err = h.service.DeleteWorkout(id)
	if err != nil {
		respondError(ctx, http.StatusInternalServerError, "could not delete workout", err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *workout) GetWorkoutExercises(ctx *gin.Context) {
	// this is where we should return only the table rows, BUT!
	// this wont be used elsewhere right now i think atleast

	id, _ := parseId(ctx)
	ex, _ := h.service.GetWorkoutExercises(id)

	setHtml200(ctx)
	template := ui.ExerciseRows(ex)
	template.Render(ctx.Request.Context(), ctx.Writer)
}
//============================ HELPERS ============================

var (
	// Errors

	// The id provided is missing or couldnt be parsed
	ErrBadID = errors.New("invalid Id")

	// Messages
	InvalidWorkoutId    = "invalid workout Id"
	InvalidInput        = "Invalid input"
	FailedToGetWorkout  = "failed to retrieve workout"
	FailedToGetWorkouts = "failed to retrieve workouts"
)

// This function returns a workout or an error
// The errors can indicate where the function encountered an error.
// ID fetching or parsing
// or fetching from the service
func (h *workout) getWorkoutByCtxId(ctx *gin.Context) (*domain.Workout, error) {
	id, err := parseId(ctx)

	if err != nil {
		return nil, err
	}

	wo, err := h.service.GetWorkout(id)
	if err != nil {
		return nil, err
	}

	return wo, nil
}

