package services

import "github.com/MagnumTrader/repforge/internal/domain"

type WorkoutService struct {
	repo domain.WorkOutRepo
}

func NewWorkoutService(repo domain.WorkOutRepo) *WorkoutService {
	return &WorkoutService{
		repo: repo,
	}
}

func (s *WorkoutService) CreateWorkout(date, kind, note string, duration int) (*domain.Workout, error) {
	workout := &domain.Workout{
		Date:     date,
		Kind:     kind,
		Duration: duration,
		Notes:    note,
	}

	if err := s.repo.SaveWorkout(workout); err != nil {
		return nil, err
	}

	return workout, nil
}

func (s *WorkoutService) DeleteWorkout(id int) error {
	return s.repo.DeleteWorkout(id)
}

func (s *WorkoutService) GetWorkout(id int) (*domain.Workout, error) {
	return s.repo.GetWorkout(id)
}

func (s *WorkoutService) GetAll() ([]domain.Workout, error) {
	return s.repo.GetAllWorkouts(0)
}
