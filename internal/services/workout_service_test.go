package services

import (

	"github.com/MagnumTrader/repforge/internal/domain"
)


type mockWRepo struct{

}

// DeleteWorkout implements domain.WorkOutRepo.
func (m mockWRepo) DeleteWorkout(id int) error {
	panic("unimplemented")
}

// GetAllWorkouts implements domain.WorkOutRepo.
func (m mockWRepo) GetAllWorkouts(userId int) ([]domain.Workout, error) {
	panic("unimplemented")
}

// GetWorkout implements domain.WorkOutRepo.
func (m mockWRepo) GetWorkout(id int) (*domain.Workout, error) {
	panic("unimplemented")
}

// SaveWorkout implements domain.WorkOutRepo.
func (m mockWRepo) SaveWorkout(workout *domain.Workout) error {
	panic("unimplemented")
}

// UpdateWorkout implements domain.WorkOutRepo.
func (m mockWRepo) UpdateWorkout(workout *domain.Workout) error {
	panic("unimplemented")
}

