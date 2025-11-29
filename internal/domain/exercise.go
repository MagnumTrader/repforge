package domain

/*
What does the exerciese containt

name, subcategory ( legs arms etc)
instructions 
link?
*/

type ExerciseRepo interface {
	GetExercise(id int) (*Exercise, error)
	GetAllExercise(userId int) ([]Exercise, error)
	SaveExercise(workout *Exercise) error
	DeleteExercise(id int) error
	UpdateExercise(workout *Exercise) error
}

type Exercise struct {
	Id int
	Name string
	Category Category
}

type Category string

const (
	CategoryLegs = "legs"
	CategoryArms = "arms"
	CategoryChest = "chest"
)



