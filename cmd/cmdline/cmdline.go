package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/MagnumTrader/repforge/internal/domain"
	"github.com/MagnumTrader/repforge/internal/infrastructure/db"
	"github.com/MagnumTrader/repforge/internal/services"
)

func main() {

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	db := db.NewDb()
	service := services.NewExerciseService(db)

	switch os.Args[1] {
	case "create":
		ex, err := parseExercise(os.Args[2:])
		if err != nil {
			slog.Error("failed to create exercise", "error", err)
			os.Exit(1)
		}
		service.CreateExercise(ex.Name, string(ex.Category))

		slog.Info("created exercise!", "exercise", ex)
		// expect object name?
		return
	case "list":
		// TODO: Take optional arguments for filters
		all, err := service.GetAll()
		if err != nil {
			slog.Error("failed to fetch exercise", "error", err)
			os.Exit(1)
		}

		for _, ex := range all {
			fmt.Printf("%+v,\n", ex)
		}
	}
}

func parseExercise(args []string) (*domain.Exercise, error) {

	if len(args) != 2 {
		slog.Info("", "len", len(args))
		return nil, fmt.Errorf("To few arguments to create exercise, required name, category")
	}

	ex := &domain.Exercise{}

	ex.Name = args[0]
	ex.Category = args[1]

	return ex, nil
}
