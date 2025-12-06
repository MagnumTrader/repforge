package domain

type WorkOutRepo interface {
	GetWorkout(id int) (*Workout, error)
	GetAllWorkouts(userId int) ([]Workout, error)
	SaveWorkout(workout *Workout) error
	DeleteWorkout(id int) error
	UpdateWorkout(workout *Workout) error
}

type Set struct {
	Id int
	Weight float32
	Reps int
}

type WorkoutExercise struct {
	Id int
	Exercise Exercise
	Sets []Set
}

type Workout struct {
	Id       int
	Date     string
	Kind     string
	Duration int
	Notes    string
	Exercises []WorkoutExercise
}

