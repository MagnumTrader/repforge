package services

import (
	"fmt"
	"log/slog"

	"github.com/MagnumTrader/repforge/internal/domain"
)

func NewExerciseService(repo domain.ExerciseRepo) *ExerciseService {
	return &ExerciseService{
		repo: repo,
	}
}

type ExerciseService struct {
	repo domain.ExerciseRepo
}

func (e *ExerciseService) CreateExercise(name, category string) (*domain.Exercise, error) {
	ex := &domain.Exercise{
		Name:     name,
		Category: domain.Category(category),
	}
	err := e.repo.SaveExercise(ex)

	if err != nil {
		return nil, fmt.Errorf("failed to save exercise: %w", err)
	}
	return ex, nil
}

func (e *ExerciseService) GetExercise(id int) (*domain.Exercise, error) {
	return e.repo.GetExercise(id)
}

func (e *ExerciseService) GetAll() ([]domain.Exercise, error) {
	all, err := e.repo.GetAllExercise(0)

	if err != nil {
		panic(err)
	}
	return all, nil
}

func (s *ExerciseService) DeleteExercise(id int) error {
	return s.repo.DeleteExercise(id)
}

func (s *ExerciseService) EditExercise(exercise *domain.Exercise) error {
	slog.Info("hello we are in the service")
	return s.repo.UpdateExercise(exercise)
}

