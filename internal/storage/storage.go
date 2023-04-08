package storage

import (
	"database/sql"
	"embed"

	"github.com/brisk84/gophkeeper/internal/config"
	"github.com/brisk84/gophkeeper/pkg/logger"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var fmigr embed.FS

type storage struct {
	lg  logger.Logger
	dsn string
	db  *sql.DB
}

func New(lg logger.Logger, cfg config.Config) (*storage, error) {
	stor := storage{
		lg:  lg,
		dsn: cfg.MainDSN,
	}
	if err := stor.Connect(); err != nil {
		return nil, err
	}
	return &stor, nil
}

// Connect Устанавливает соединение
func (s *storage) Connect() error {
	var err error
	if s.db, err = sql.Open("postgres", s.dsn); err != nil {
		return err
	}
	if err = s.db.Ping(); err != nil {
		return err
	}
	goose.SetBaseFS(fmigr)
	if err = goose.Up(s.db, "migrations"); err != nil {
		return err
	}
	return nil
}

// Close Закрывает соединение
func (s *storage) Close() error {
	return s.db.Close()
}
