package db

import (
	"fmt"
	"github.com/MagnumTrader/repforge/internal/domain"
)

func NewInMem() *InMem {
	return &InMem{}
}

type InMem struct{}

func (d *InMem) GetWorkout(id int) (*domain.Workout, error) {
	for _, wo := range workouts {
		if wo.Id == id {
			return &wo, nil
		}
	}
	return nil, fmt.Errorf("Workout with id %d not found!", id)
}
func (d *InMem) GetAllWorkouts(userId int) ([]domain.Workout, error) {
	return workouts, nil
}

// TODO: Implement this
func (d *InMem) SaveWorkout(workout domain.Workout) error {
	panic("Not implemented")
}

var workouts = []domain.Workout{
	{
		Id:       1,
		Date:     "2025-11-10",
		Kind:     "Running",
		Duration: 30,
		Notes:    "Morning jog in the park",
	},
	{
		Id:       2,
		Date:     "2025-11-09",
		Kind:     "Cycling",
		Duration: 45,
		Notes:    "Evening ride with friends",
	},
	{
		Id:       3,
		Date:     "2026-11-08",
		Kind:     "Yoga",
		Duration: 60,
		Notes:    "Relaxing session at home",
	},
	{
		Id:       4,
		Date:     "2026-11-08",
		Kind:     "Gym",
		Duration: 45,
		Notes:    "",
	},
}

var exercise = []domain.Exercise {
	{
		Id:   1,
		Name: "Test Exercise",
	},
}

// Exercise repo impl
func (d *InMem) GetExercise(id int) (*domain.Exercise, error) {

	if id != 1 {
		panic("no exercise like that")
	}

	return &exercise[0], nil
	
}
func (d *InMem) GetAllExercise(userId int) ([]domain.Exercise, error) {
	return exercise, nil
}
func (d *InMem) SaveExercise(exercise *domain.Exercise) error {
	panic("not implemented")
}
func (d *InMem) DeleteExercise(id int) error {
	panic("not implemented")
}
func (d *InMem) UpdateExercise(workout *domain.Exercise) error {
	panic("not implemented")
}
