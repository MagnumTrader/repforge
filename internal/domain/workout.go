package domain

type WorkOutRepo interface {
	GetWorkout(id int) (*Workout, error)
	GetAllWorkouts(userId int) ([]Workout, error)
	SaveWorkout(workout Workout) error
}

type Workout struct {
	Id       int
	Date     string
	Type     string
	Duration int
	Notes    string
}

