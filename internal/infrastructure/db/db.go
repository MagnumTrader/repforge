package db

import (
	"database/sql"

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

	mo := &domain.Workout{
		Id:       id,
		Date:     "",
		Kind:     "",
		Duration: 0,
		Notes:    "",
	}

	row.Scan(
		&mo.Id,
		&mo.Date,
		&mo.Kind,
		&mo.Duration,
		&mo.Notes,
	)
	return mo, nil
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
	_, err := d.inner.Exec("delete from workouts where id = ?", id)
	return err
}
