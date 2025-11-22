package domain

type WorkOutRepo interface {
	GetWorkout(id int) (*Workout, error)
	GetAllWorkouts(userId int) ([]Workout, error)
	SaveWorkout(workout *Workout) error
	DeleteWorkout(id int) error
	UpdateWorkout(workout *Workout) error
}


type Workout struct {
	Id       int
	Date     string
	Kind     string
	Duration int
	Notes    string
}

