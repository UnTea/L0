package migrations

import (
	"database/sql"
	"github.com/UnTea/L0/internal/config"
	"github.com/UnTea/L0/pkg/helpers"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

// MakeMigrations - makes all migration via goose up
func MakeMigrations(dsn string, cfg config.Config) error {
	mdb, _ := sql.Open("postgres", dsn)
	err := mdb.Ping()
	if err != nil {
		return err
	}
	defer helpers.Closer(mdb)

	err = goose.Up(mdb, cfg.Database.Migrations)
	if err != nil {
		return err
	}
	return nil
}
