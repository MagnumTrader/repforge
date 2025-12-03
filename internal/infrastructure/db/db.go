package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/MagnumTrader/repforge/internal/domain"
	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	inner *sql.DB
}

func NewDb() *Db {
	db, err := sql.Open("sqlite3", "data/repforge.db")
	if err != nil {
		panic(err)
	}

	return &Db{
		inner: db,
	}
}

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

	workouts = make([]domain.Workout, 0)

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

func (d *Db) SaveWorkout(workout *domain.Workout) error {
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

// DeleteExercise implements domain.ExerciseRepo.
func (d *Db) DeleteExercise(id int) error {
	panic("unimplemented")
}

// GetAllExercise implements domain.ExerciseRepo.
func (d *Db) GetAllExercise(userId int) ([]domain.Exercise, error) {
	panic("unimplemented")
}

// GetExercise implements domain.ExerciseRepo.
func (d *Db) GetExercise(id int) (*domain.Exercise, error) {
	panic("unimplemented")
}

// SaveExercise implements domain.ExerciseRepo.
func (d *Db) SaveExercise(workout *domain.Exercise) error {

	query := "insert into exercises (name, category) values (?, ?)"

	result, err := d.inner.Exec(query, workout.Name, workout.Category)
	if err != nil {
		slog.Error("Failed to insert exercise", "error", err)
		return err
	}

	// NOTE: fails only if not supported (i think)
	id, _ := result.LastInsertId()

	workout.Id = int(id)

	return nil
}

// UpdateExercise implements domain.ExerciseRepo.
func (d *Db) UpdateExercise(workout *domain.Exercise) error {
	panic("unimplemented")
}
