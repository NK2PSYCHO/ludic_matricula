package testhelper

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	_ "github.com/uptrace/bun/driver/pgdriver"
)

func NewTestDB(t *testing.T) *bun.DB {
	t.Helper()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	sqldb, err := sql.Open("pg", dsn)
	if err != nil {
		t.Fatalf("open test db: %v", err)
	}

	if err := sqldb.Ping(); err != nil {
		t.Fatalf("ping test db: %v — is postgres_test_db running?", err)
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	t.Cleanup(func() { sqldb.Close() })

	return db
}

func TruncateTables(t *testing.T, db *bun.DB, tables ...string) {
	t.Helper()
	for _, table := range tables {
		if _, err := db.Exec("TRUNCATE TABLE ? RESTART IDENTITY CASCADE", bun.Ident(table)); err != nil {
			t.Fatalf("truncate %s: %v", table, err)
		}
	}
}
