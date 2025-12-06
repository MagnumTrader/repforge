package domain

type WorkOutRepo interface {
	GetWorkout(id int) (*Workout, error)
	GetAllWorkouts(userId int) ([]Workout, error)
	CreateWorkout(workout *Workout) error
	DeleteWorkout(id int) error
	UpdateWorkout(workout *Workout) error
}

type WorkoutExerciseRepo interface {
	// Todo: should this be an id to a workout, not id of a workout exercise?
	GetWorkoutExercise(id int) (*WorkoutExercise, error)
	GetAllWorkoutExercises(userId int) ([]WorkoutExercise, error)
	GetAllForWorkout(workout_id int) ([]WorkoutExercise, error)
	CreateWorkoutExercise(workoutExercise *WorkoutExercise) error
	DeleteWorkoutExercise(id int) error
	UpdateWorkoutExercise(workoutExercise *WorkoutExercise) error
}

type Set struct {
	Id     int
	Weight float32
	Reps   int
}

type WorkoutExercise struct {
	Id       int
	Exercise Exercise
	Sets     []Set
}

type Workout struct {
	Id        int
	Date      string
	Kind      string
	Duration  int
	Notes     string
	Exercises []WorkoutExercise
}
