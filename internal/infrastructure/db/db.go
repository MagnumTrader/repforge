package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"

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

	database := Db{
		inner: db,
	}


	if err := database.runMigrations(); err != nil {
		slog.Error("Failed to run migrations", "error", err)
		os.Exit(1)
	}

	return &database
}

//go:embed migrations/*.sql
var migrationsFS embed.FS

func (db *Db) runMigrations() error {
    // Ensure migrations table exists
    _, err := db.inner.Exec(`
        CREATE TABLE IF NOT EXISTS schema_migrations (
            version INTEGER PRIMARY KEY,
            applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        return err
    }

    // Get applied migrations
    appliedMigrations, err := db.getAppliedMigrations()
    if err != nil {
        return err
    }

    // Read migration files
    entries, err := migrationsFS.ReadDir("migrations")
    if err != nil {
        return err
    }

    // Sort migrations by filename
    sort.Slice(entries, func(i, j int) bool {
        return entries[i].Name() < entries[j].Name()
    })

    // Run pending migrations
    for _, entry := range entries {
        if !strings.HasSuffix(entry.Name(), ".sql") {
            continue
        }

        version := extractVersion(entry.Name())
        if appliedMigrations[version] {
            continue // Already applied
        }

        slog.Info("Running migration", "file", entry.Name())
        
        content, err := migrationsFS.ReadFile("migrations/" + entry.Name())
        if err != nil {
            return err
        }

        // Run migration in transaction
        tx, err := db.inner.Begin()
        if err != nil {
            return err
        }

        if _, err := tx.Exec(string(content)); err != nil {
            tx.Rollback()
            return fmt.Errorf("migration %s failed: %w", entry.Name(), err)
        }

        // Record migration
        _, err = tx.Exec("INSERT INTO schema_migrations (version) VALUES (?)", version)
        if err != nil {
            tx.Rollback()
            return err
        }

        if err := tx.Commit(); err != nil {
            return err
        }

        slog.Info("Migration %s completed", "file", entry.Name())
    }

    return nil
}

func (db *Db) getAppliedMigrations() (map[int]bool, error) {
    rows, err := db.inner.Query("SELECT version FROM schema_migrations")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    applied := make(map[int]bool)
	var version int

    for rows.Next() {
        if err := rows.Scan(&version); err != nil {
            return nil, err
        }
        applied[version] = true
    }

    return applied, nil
}

func extractVersion(filename string) int {
    // Extract number from "001_init.sql" -> 1
    var version int
    fmt.Sscanf(filename, "%d", &version)
    return version
}

