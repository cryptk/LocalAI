package database

import (
	"database/sql"
	"embed"
	"os"
	"path/filepath"

	"github.com/go-skynet/LocalAI/core/config"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/rs/zerolog/log"
)

type DB struct {
	DB *Queries
}

//go:embed migrations/*
var dbMigrations embed.FS

func Open(appConfig *config.ApplicationConfig) (*DB, error) {
	os.MkdirAll(filepath.Dir(appConfig.DBPath), 0755)

	database, err := sql.Open("sqlite3", appConfig.DBPath)
	if err != nil {
		return nil, err
	}
	err = database.Ping()
	if err != nil {
		return nil, err
	}
	log.Debug().Msg("database connection established")

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations/sqlite3",
	}

	n, err := migrate.Exec(database, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return nil, err
	}
	log.Debug().Int("numMigrations", n).Msg("applied database migrations")

	db := &DB{
		DB: New(database),
	}

	return db, nil
}
