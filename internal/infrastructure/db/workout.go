package db

import (
	"fmt"
	"log/slog"

	"github.com/MagnumTrader/repforge/internal/domain"
)


func (d *Db) GetWorkout(id int) (*domain.Workout, error) {
	row := d.inner.QueryRow("select id, date, type, duration, notes  from workouts where id=?", id)

	wo := &domain.Workout{
		Id:       id,
		Date:     "",
		Kind:     "",
		Duration: 0,
		Notes:    "",
	}

	row.Scan(
		&wo.Id,
		&wo.Date,
		&wo.Kind,
		&wo.Duration,
		&wo.Notes,
	)
	return wo, nil
}

func (d *Db) GetAllWorkouts(userId int) ([]domain.Workout, error) {
	rows, err := d.inner.Query("select id, date, type, duration, notes  from workouts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workouts := make([]domain.Workout, 0)

	wo := &domain.Workout{}

	for rows.Next() {
		err := rows.Scan(
			&wo.Id,
			&wo.Date,
			&wo.Kind,
			&wo.Duration,
			&wo.Notes)

		if err != nil {
			return nil, err
		}

		workouts = append(workouts, *wo)

	}

	return workouts, nil
}

func (d *Db) CreateWorkout(workout *domain.Workout) error {
	row, err := d.inner.Exec("INSERT INTO workouts (date, duration, type, notes) VALUES (?, ?, ?, ?)",
		workout.Date,
		workout.Duration,
		workout.Kind,
		workout.Notes,
	)

	if err != nil {
		return err
	}

	// Sqlite3 supports this and we have autoincrement on The id
	id, _ := row.LastInsertId()
	workout.Id = int(id)

	return nil
}

func (d *Db) DeleteWorkout(id int) error {
	res, err := d.inner.Exec("delete from workouts where id = ?", id)
	rows, err := res.RowsAffected()
	// TODO: should return error when not found
	if rows == 0 {
		return fmt.Errorf("Zero records was updated")
	}
	slog.Info("delete request successfull", "Deleted rows", rows, "workout id", id)
	return err
}

// UpdateWorkout implements domain.WorkOutRepo.
func (d *Db) UpdateWorkout(workout *domain.Workout) error {
	query := "UPDATE workouts SET date = ?, duration = ?, type = ?, notes = ? where id = ?"
	rows, err := d.inner.Exec(query,
		&workout.Date,
		&workout.Duration,
		&workout.Kind,
		&workout.Notes,
		&workout.Id,
	)

	// NOTE: Only errors if db does not support rowsaffected (i think)
	affected, _ := rows.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("Zero records was updated")
	}

	return err
}



// Create implements domain.CrudRepo.
func (d *Db) CreateWorkoutExercise(workoutId, exerciseId int) error {
	query := `insert into workoutexercises
			(workout_id, exercise_id, order_in_exercise) 
			select ?, ?, coalesce(max(order_in_exercise), 0) + 1 
			from workoutexercises 
			where workout_id = ?)`

	res, err := d.inner.Exec(query, workoutId, exerciseId, workoutId)

	if err != nil {
		return fmt.Errorf("Failed to create workoutexercise error: %w", err)
	}

	affected, err := res.RowsAffected()

	if affected <= 0 && err == nil {
		// No error but affected is still 0, something is wrong
		return fmt.Errorf("No changes, and no error")
	}
	return nil
}

// Delete implements domain.CrudRepo.
func (w *Db) DeleteWorkoutExercise(id int) error {
	panic("unimplemented")
}

// Get implements domain.CrudRepo.
func (w *Db) GetWorkoutExercise(id int) (*domain.WorkoutExercise, error) {
	panic("unimplemented")
}

func (d *Db) GetAllForWorkout(workout_id int) ([]domain.WorkoutExercise, error) {
	query := `select w.id, e.id as eId, e.name, e.category from workoutexercises w left join exercises e on w.exercise_id=e.id where workout_id = ?`
	rows, err := d.inner.Query(query, workout_id)
	if err != nil {
	  return nil, err
	}
	defer rows.Close()

	wos := make([]domain.WorkoutExercise, 0)
	var wId, eId int
	var name, category string
	for rows.Next() {

		err := rows.Scan(&wId, &eId, &name, &category)
		if err != nil {
			slog.Error("failed to scan workoutexercise", "id", eId, "name", name, "category", category)
			continue
		}

		wo := domain.WorkoutExercise{
			Id:       wId,
			Exercise: domain.Exercise{
				Id:       eId,
				Name:     name,
				Category: category,
			},
			Sets:     []domain.Set{},
		}
		wos = append(wos, wo)
	}
	return wos, nil
}

// GetAll implements domain.CrudRepo.
func (w *Db) GetAllWorkoutExercises(userId int) ([]domain.WorkoutExercise, error) {
	panic("unimplemented")
}

// Update implements domain.CrudRepo.
func (w *Db) UpdateWorkoutExercise(workoutexercise *domain.WorkoutExercise) error {
	panic("unimplemented")
}

func (d *Db) AddSet(WeId int, set *domain.Set) error {return nil}
func (d *Db) DeleteSet(id int) error {return nil}
func (d *Db) UpdateSet(set *domain.Set) error {return nil}
