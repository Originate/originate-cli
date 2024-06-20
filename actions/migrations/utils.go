package migrations

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func connectToDatabase() (*sql.DB, error) {
	dsn := os.Getenv("ORIGINATE_DATABASE_DSN")
	if dsn == "" {
		return &sql.DB{}, errors.New("ORIGINATE_DATABASE_DSN environment variable isn't set")
	}

	return sql.Open("postgres", dsn)
}

func createPostgresProvider(migrationsDir string) (*goose.Provider, error) {
	db, err := connectToDatabase()
	if err != nil {
		return &goose.Provider{}, err
	}

	return goose.NewProvider(
		goose.DialectPostgres,
		db,
		os.DirFS(migrationsDir),
	)
}
