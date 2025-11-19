package routes

import (
	"github.com/MagnumTrader/repforge/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterWorkoutRoutes(r *gin.Engine, service services.WorkoutService) {
	handler := handlers.NewWorkoutHandler(service)

	grp := r.Group("/workouts")
	grp.GET("/", handler.WorkoutsListHandler)

}

