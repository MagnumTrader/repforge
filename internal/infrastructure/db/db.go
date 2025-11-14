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

// So this

func (d *Db) GetWorkout(id int) (*domain.Workout, error) {
	row := d.inner.QueryRow("select id, date, type, duration, notes  from workouts where id=?", id)

	mo := &domain.Workout{
		Id:       id,
		Date:     "",
		Type:     "",
		Duration: 0,
		Notes:    "",
	}

	row.Scan(
		&mo.Id,
		&mo.Date,
		&mo.Type,
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
			&wo.Type,
			&wo.Duration,
			&wo.Notes)

		if err != nil {
			return nil, err
		}

		workouts = append(workouts, *wo)

	}

	return workouts, nil
}
func (d *Db) SaveWorkout(workout domain.Workout) error {
	return nil
}
