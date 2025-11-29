package services

import "github.com/MagnumTrader/repforge/internal/domain"


func NewExerciseService(repo domain.ExerciseRepo) *ExerciseService  {
	return &ExerciseService{
		repo: repo,
	}
}

type ExerciseService struct {
	repo domain.ExerciseRepo
}

func (e *ExerciseService)GetExercise(id int) (*domain.Exercise, error)  {
	return e.repo.GetExercise(id)
}

func (e *ExerciseService)GetAll() ([]domain.Exercise, error)  {
	all, err :=  e.repo.GetAllExercise(0) 

	if err != nil {
	  panic(err)
	}
	return all, nil
}
