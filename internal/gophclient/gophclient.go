package gophclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/brisk84/gophkeeper/api"
	"github.com/brisk84/gophkeeper/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GophClient struct {
	creds   credentials.TransportCredentials
	srvAddr string
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}

func New(srvAddr string) (*GophClient, error) {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		return nil, err
	}
	return &GophClient{creds: tlsCredentials, srvAddr: srvAddr}, nil
}

func (g *GophClient) Register(user domain.User) (string, error) {
	conn, err := grpc.Dial(g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := cl.Register(ctx, &api.RegisterLoginReq{
		Login:    user.Login,
		Password: user.Password,
	})
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func (g *GophClient) Login(user domain.User) (string, error) {
	conn, err := grpc.Dial(g.srvAddr, grpc.WithTransportCredentials(g.creds))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cl := api.NewGophKeeperClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := cl.Login(ctx, &api.RegisterLoginReq{
		Login:    user.Login,
		Password: user.Password,
	})
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}
