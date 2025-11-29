package handlers

import (
	"fmt"
	"net/http"

	"github.com/MagnumTrader/repforge/internal/http/ui"
	"github.com/MagnumTrader/repforge/internal/services"
	"github.com/gin-gonic/gin"
)

func NewExerciseHandler(service *services.ExerciseService) *exerciseHandler {
	handle := &exerciseHandler{
		service: service,
	}
	return handle
}

type exerciseHandler struct {
	service *services.ExerciseService
}

func (e *exerciseHandler) ExerciseDetails(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		respondError(c, http.StatusBadRequest, "Invalid Exercise Id", err)
		return
	}
	ex, err := e.service.GetExercise(id)
	if err != nil {
		respondError(c, http.StatusInternalServerError, "Failed to retrieve Exercise", err)
		return
	}
	// TODO: Render templates for this shiet
	c.String(http.StatusOK, fmt.Sprintf("exercise Name: %s, category: %s", ex.Name, ex.Category))
}

func (e *exerciseHandler) ExerciseList(c *gin.Context) {

	all, err := e.service.GetAll()

	if err != nil {
		respondError(c, http.StatusInternalServerError, "Failed to fetch exercises", err)
		return
	}

	setHtml200(c)

	template := ui.ExerciseListPartial(all)
	if !IsHtmxRequest(c) {
		template = ui.Base(template)
	}
	template.Render(c.Request.Context(), c.Writer)
}

func (e *exerciseHandler) NewExercise(c *gin.Context) {

	// parse the exercise from the request
	ex := struct {
		Name     string `form:"name" binding:"required"`
		Category string `form:"category" binding:"required"`
	}{}

	err := c.ShouldBind(&ex)
	if err != nil {
		respondError(c, http.StatusBadRequest, "Failed to parse Exercise from form", err)
		return
	}

	e.service.CreateExercise(ex.Name, ex.Category)
}
