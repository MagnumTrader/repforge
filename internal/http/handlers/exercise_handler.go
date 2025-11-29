package handlers

import (
	"github.com/MagnumTrader/repforge/internal/services"
	"github.com/gin-gonic/gin"
)

type exerciseHandler struct {
	service *services.ExerciseService
}

func (e *exerciseHandler) ExerciseList(c *gin.Context) {
	exercise, _ := e.service.GetExercise(1)
	c.Writer.Write([]byte(exercise.Name))
}

func NewExerciseHandler(service *services.ExerciseService) *exerciseHandler {
	handle := &exerciseHandler{
		service: service,
	}


	return handle

}
