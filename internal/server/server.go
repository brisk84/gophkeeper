package server

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/brisk84/gophkeeper/api"
	"github.com/brisk84/gophkeeper/internal/config"
	"github.com/brisk84/gophkeeper/internal/handler"
	"github.com/brisk84/gophkeeper/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	srv  *grpc.Server
	lg   logger.Logger
	addr string
	h    *handler.Handler
}

func New(lg logger.Logger, cfg config.Config, h *handler.Handler) (*Server, error) {
	var err error
	srv := &Server{
		lg:   lg,
		addr: cfg.AppAddr,
		h:    h,
	}
	loggingInterceptor := grpc.ChainUnaryInterceptor(srv.loggingMiddleware)
	creds, err := loadTLSCredentials()
	if err != nil {
		return nil, err
	}
	srv.srv = grpc.NewServer(grpc.Creds(creds), loggingInterceptor)
	api.RegisterGophKeeperServer(srv.srv, srv.h)
	return srv, nil
}

func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		s.lg.Errorln(err.Error())
		return err
	}
	go func() {
		<-ctx.Done()
		s.srv.GracefulStop()
	}()
	return s.srv.Serve(lis)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}
	return credentials.NewTLS(config), nil
}
