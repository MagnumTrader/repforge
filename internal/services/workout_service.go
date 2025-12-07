package services

import "github.com/MagnumTrader/repforge/internal/domain"

type WorkoutService struct {
	workoutRepo         domain.WorkOutRepo
	workoutExerciseRepo domain.WorkoutExerciseRepo
}

func NewWorkoutService(woRepo domain.WorkOutRepo, woExRepo domain.WorkoutExerciseRepo) *WorkoutService {
	return &WorkoutService{
		workoutRepo:         woRepo,
		workoutExerciseRepo: woExRepo,
	}
}

func (s *WorkoutService) CreateWorkout(date, kind, note string, duration int) (*domain.Workout, error) {
	workout := &domain.Workout{
		Date:     date,
		Kind:     kind,
		Duration: duration,
		Notes:    note,
	}

	if err := s.workoutRepo.CreateWorkout(workout); err != nil {
		return nil, err
	}

	return workout, nil
}

func (s *WorkoutService) EditWorkout(workout *domain.Workout) error {
	return s.workoutRepo.UpdateWorkout(workout)
}
func (s *WorkoutService) DeleteWorkout(id int) error {
	return s.workoutRepo.DeleteWorkout(id)
}

func (s *WorkoutService) GetWorkout(id int) (*domain.Workout, error) {
	return s.workoutRepo.GetWorkout(id)
}

func (s *WorkoutService) GetAll() ([]domain.Workout, error) {
	return s.workoutRepo.GetAllWorkouts(0)
}

func (s *WorkoutService) GetWorkoutExercises(workout_id int) ([]domain.WorkoutExercise, error) {
	// we are on the details page, and we need all the workoutsexercises for this workout
	// the workout service already have the workoutExerciseRepo
	// the template should return a partial that then calls into what it should have
	all, err := s.workoutExerciseRepo.GetAllForWorkout(workout_id)
	if err != nil {
	  return nil, err
	}
	
	return all, nil
}
