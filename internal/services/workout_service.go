package service

import "github.com/MagnumTrader/repforge/internal/domain"

type WorkoutService struct {
	repo domain.WorkOutRepo
}

func NewWorkoutService(repo domain.WorkOutRepo) WorkoutService {
	return WorkoutService{
		repo: repo,
	}
}

func (s *WorkoutService)CreateWorkout(name string) error {
	return nil
}
func (s *WorkoutService)GetAll() ([]domain.Workout, error) {
	return nil, nil
}


