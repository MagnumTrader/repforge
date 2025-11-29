package db

import (
	"fmt"
	"slices"

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

var exercises = []domain.Exercise {
	{
		Id:   1,
		Name: "Leg Press",
		Category: domain.CategoryLegs,
	},
	{
		Id:   2,
		Name: "Bench Press",
		Category: domain.CategoryChest,
	},
}

// Exercise repo impl
func (d *InMem) GetExercise(id int) (*domain.Exercise, error) {
	for _, ex := range exercises {
		if ex.Id == id {
			return &ex, nil
		}
	}

	return nil, fmt.Errorf("No exercise with id %d", id)
}

func (d *InMem) GetAllExercise(userId int) ([]domain.Exercise, error) {
	return exercises, nil
}

func (d *InMem) SaveExercise(exercise *domain.Exercise) error {
	var maxId int
	for _, e := range exercises {
		maxId = max(maxId, e.Id)
	}
	exercise.Id = maxId + 1
	exercises = append(exercises, *exercise)
	return nil
}

func (d *InMem) DeleteExercise(id int) error {
	for i, e := range exercises {
		if e.Id == id {
			exercises = slices.Delete(exercises, i , i + 1)
			return nil
		}
	}
	return fmt.Errorf("Unable to delete exercise with id %d", id)
}

func (d *InMem) UpdateExercise(exercise *domain.Exercise) error {
	panic("not implemented")
}
