package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/brisk84/gophkeeper/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

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

func intTests(ctx context.Context, c api.GophKeeperClient) int {
	resp, err := c.Register(ctx, &api.RegisterLoginReq{
		Login:    "Vasya",
		Password: "123",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Token)
	return 0
}

func main() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := api.NewGophKeeperClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	exitCode := intTests(ctx, c)
	fmt.Println("ExitCode:", exitCode)

	cancel()
	cc1.Close()
	os.Exit(exitCode)
}
