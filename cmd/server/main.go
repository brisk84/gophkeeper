package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/brisk84/gophkeeper/internal/app"
	"github.com/brisk84/gophkeeper/internal/config"
	"github.com/brisk84/gophkeeper/pkg/logger"
)

func main() {
	lg, err := logger.New(true)
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.New()
	if err != nil {
		lg.Fatal(err)
	}
	a, err := app.New(lg, cfg)
	if err != nil {
		lg.Fatal(err)
	}
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()
	if err := a.Run(ctx); err != nil {
		lg.Fatalln(err)
	}
}
