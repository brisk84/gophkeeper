package app

import (
	"context"

	"github.com/brisk84/gophkeeper/internal/config"
	"github.com/brisk84/gophkeeper/internal/handler"
	"github.com/brisk84/gophkeeper/internal/server"
	"github.com/brisk84/gophkeeper/internal/storage"
	"github.com/brisk84/gophkeeper/internal/usecase"
	"github.com/brisk84/gophkeeper/pkg/logger"
)

type App struct {
	srv *server.Server
	lg  logger.Logger
}

func New(lg logger.Logger, cfg config.Config) (*App, error) {
	stor, err := storage.New(lg, cfg)
	if err != nil {
		return nil, err
	}
	uc, err := usecase.New(lg, stor)
	if err != nil {
		return nil, err
	}
	h := handler.New(lg, uc)
	srv, err := server.New(lg, cfg, h)
	if err != nil {
		return nil, err
	}
	return &App{
		srv: srv,
		lg:  lg,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.srv.Start(ctx)
}
