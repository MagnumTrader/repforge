package routes

import (
	"github.com/MagnumTrader/repforge/internal/http/handlers"
	"github.com/MagnumTrader/repforge/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterWorkoutRoutes(r *gin.Engine, service *services.WorkoutService) {
	handler := handlers.NewWorkoutHandler(service)

	grp := r.Group("/workouts")

	grp.GET("", handler.WorkoutsList)
	grp.GET("/:id", handler.WorkoutDetails)
	grp.GET("/:id/edit", handler.EditWorkoutForm) // this should return the form
	grp.PATCH("/:id", handler.WorkoutDetails) // form send here
	grp.DELETE("/:id", handler.DeleteWorkout)
	grp.GET("/new", handler.NewWorkoutForm)
	grp.POST("/new", handler.NewWorkout)
}

