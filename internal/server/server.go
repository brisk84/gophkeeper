package server

import (
	"context"
	"net/http"
	"time"

	"github.com/brisk84/gophkeeper/internal/config"
	"github.com/brisk84/gophkeeper/internal/handler"
	"github.com/brisk84/gophkeeper/internal/router"
	"github.com/brisk84/gophkeeper/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	lg  logger.Logger
}

func New(lg logger.Logger, cfg config.Config, h *handler.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Handler: router.New(h),
			Addr:    cfg.AppAddr,
		},
		lg: lg,
	}
}

func (s *Server) Start(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(s.srv.ListenAndServe)
	eg.Go(func() error {
		s.lg.Infoln("GophKeeper started. App addr:", s.srv.Addr)
		<-ctx.Done()
		s.lg.Infoln("GophKeeper shutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := s.srv.Shutdown(shutdownCtx)
		if err != nil {
			s.lg.Errorln(err.Error())
		}
		return err
	})
	return eg.Wait()
}
