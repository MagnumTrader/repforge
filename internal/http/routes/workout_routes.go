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
	grp.PUT("/:id", handler.EditWorkout)
	grp.GET("/edit/:id", handler.EditWorkoutForm)
	grp.GET("/new", handler.NewWorkoutForm)
	grp.POST("/new", handler.NewWorkout)
	grp.DELETE("/:id", handler.DeleteWorkout)
}

