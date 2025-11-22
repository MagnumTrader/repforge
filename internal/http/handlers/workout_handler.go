package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

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

	if !IsHtmxRequest(ctx) {
		template = ui.Base(template)
	}
	template.Render(ctx.Request.Context(), ctx.Writer)
}

func (h *workout) WorkoutDetails(ctx *gin.Context) {
	workout, err := h.getWorkoutFromCtx(ctx)
	if err != nil {
		if errors.Is(err, ErrBadID) {
			respondError(ctx, http.StatusBadRequest, InvalidWorkoutId, err)
			return
		}
		respondError(ctx, http.StatusInternalServerError, FailedToGetWorkout, err)
		return
	}

	template := ui.WorkoutDetailsPartial(*workout)
	if !IsHtmxRequest(ctx) {
		template = ui.Base(template)
	}
	template.Render(ctx.Request.Context(), ctx.Writer)
}

func (h *workout) EditWorkoutForm(ctx *gin.Context) {
	wo, err := h.getWorkoutFromCtx(ctx)

	if err != nil {
		if errors.Is(err, ErrBadID) {
			respondError(ctx, http.StatusBadRequest, InvalidWorkoutId, err)
			return
		}
		respondError(ctx, http.StatusInternalServerError, FailedToGetWorkout, err)
		return
	}

	if IsHtmxRequest(ctx) {
		// We should render the partial
		template := ui.WorkoutForm(wo)
		template.Render(ctx.Request.Context(), ctx.Writer)
		return
	}

	// NOTE: maybe support standalone screen later
	ctx.Status(http.StatusNotFound)
}

func (h *workout) NewWorkoutForm(ctx *gin.Context) {
	if IsHtmxRequest(ctx) {
		// we should render the partial
		template := ui.WorkoutForm(nil)
		template.Render(ctx.Request.Context(), ctx.Writer)
		return
	}

	// NOTE: maybe support standalone screen later
	ctx.Status(http.StatusNotFound)
}

func (h *workout) DeleteWorkout(ctx *gin.Context) {

	id, err := h.parseId(ctx)
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
		respondError(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	workout, err := h.service.CreateWorkout(wo.Date, wo.Type, wo.Note, wo.Duration)

	if err != nil {
		respondError(ctx, http.StatusInternalServerError, "Invalid input", err)
		slog.Error("Failed to insert workout into DB", "Error:", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	row := ui.WorkoutTableRow(*workout)
	row.Render(ctx.Request.Context(), ctx.Writer)
}

//============================ HELPERS ============================

var (
	// Errors

	// The id provided is missing or couldnt be parsed
	ErrBadID         = errors.New("invalid Id")

	// Messages
	InvalidWorkoutId = "invalid workout Id"
	FailedToGetWorkout = "failed to retrieve workout"
	FailedToGetWorkouts = "failed to retrieve workouts"
)

// This function returns a workout or an error
// The errors can indicate where the function encountered an error.
// ID fetching or parsing
// or fetching from the service
func (h *workout) getWorkoutFromCtx(ctx *gin.Context) (*domain.Workout, error) {
	id, err := h.parseId(ctx)

	if err != nil {
		return nil, err
	}

	wo, err := h.service.GetWorkout(id)
	if err != nil {
		return nil, err
	}

	return wo, nil
}

// Parse the id of the request, MAY be more general later if we have structure for system wide Id's
func (h *workout) parseId(ctx *gin.Context) (int, error) {
	idString := ctx.Param("id")
	if idString == "" {
		return 0, fmt.Errorf("%w: missing", ErrBadID)
	}

	id, err := strconv.Atoi(idString)

	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrBadID, idString)
	}
	return id, nil
}
