package db

import (
	"fmt"
	"log/slog"

	"github.com/MagnumTrader/repforge/internal/domain"
)


const exerciseDbName = "exercises"

// DeleteExercise implements domain.ExerciseRepo.
func (d *Db) DeleteExercise(id int) error {

	query := fmt.Sprintf("DELETE from %s where id = %d", exerciseDbName, id)

	res, err := d.inner.Exec(query)

	if err != nil {
		return fmt.Errorf("Failed to delete exercise with id %d: %w", id, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		// not supported so cant check rows affected
	  return nil
	}

	if affected == 0 {
		return fmt.Errorf("Failed to delete exercise, 0 rows affected!")
	}

	return nil
}

// GetAllExercise implements domain.ExerciseRepo.
func (d *Db) GetAllExercise(userId int) ([]domain.Exercise, error) {

	query := fmt.Sprintf("select * from %s", exerciseDbName)
	rows, err := d.inner.Query(query)

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch all exercises %w", err)
	}
		
	exerciseList := []domain.Exercise{}
	for rows.Next() {
		ex := domain.Exercise{}
		rows.Scan(&ex.Id, &ex.Name, &ex.Category)
		exerciseList = append(exerciseList, ex)
	}

	return exerciseList, nil
}

func (d *Db) GetExercise(id int) (*domain.Exercise, error) {
	panic("unimplemented")
}

// SaveExercise implements domain.ExerciseRepo.
func (d *Db) CreateExercise(exercise *domain.Exercise) error {

	query := fmt.Sprintf("insert into %s (name, category) values (?, ?)", exerciseDbName)

	result, err := d.inner.Exec(query, exercise.Name, exercise.Category)
	if err != nil {
		slog.Error("Failed to insert exercise", "error", err)
		return err
	}

	// NOTE: fails only if not supported (i think)
	id, _ := result.LastInsertId()

	exercise.Id = int(id)

	return nil
}

// UpdateExercise implements domain.ExerciseRepo.
func (d *Db) UpdateExercise(ex *domain.Exercise) error {
	query := "UPDATE %s set name=?, category=? where id = ?"
	if _, err := d.inner.Exec(query, ex.Name, ex.Category, ex.Id); err != nil {
		return err
	}
	return nil
}



