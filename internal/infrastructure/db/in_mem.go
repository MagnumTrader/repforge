package db

import (
	"fmt"

	"github.com/MagnumTrader/repforge/internal/domain"
)

type InMem struct{}

func (d *InMem) GetWorkout(id int) (*domain.Workout, error) {
	for _, wo := range domain.Workouts {
		if wo.Id == id {
			return &wo, nil
		}
	}
	return nil, fmt.Errorf("Workout with id %d not found!", id)
}
func (d *InMem) GetAllWorkouts(userId int) ([]domain.Workout, error) {
	return domain.Workouts, nil
}
// TODO: Implement this
func (d *InMem) SaveWorkout(workout domain.Workout) error {
	return nil
}
