package migrations

import (
	"database/sql"
	"github.com/UnTea/L0/internal/config"
	"github.com/UnTea/L0/pkg/helpers"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func MakeMigrations(connectionString string, cfg config.Config) error {
	mdb, _ := sql.Open("postgres", connectionString)

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
