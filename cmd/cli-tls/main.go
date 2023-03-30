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

	pb "github.com/brisk84/gophkeeper/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func intTests(ctx context.Context, c pb.GreeterClient) int {
	resp, err := c.SayHello(ctx, &pb.HelloRequest{
		Name: "Vasya",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
	return 0

}

func main() {
	// conn, err := grpc.Dial("localhost:4343", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	cc1, err := grpc.Dial("0.0.0.0:4343", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	// cc2, err := grpc.Dial(
	// 	*serverAddress,
	// 	grpc.WithTransportCredentials(tlsCredentials),
	// 	grpc.WithUnaryInterceptor(interceptor.Unary()),
	// 	grpc.WithStreamInterceptor(interceptor.Stream()),
	// )
	// if err != nil {
	// 	log.Fatal("cannot dial server: ", err)
	// }

	c := pb.NewGreeterClient(cc1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	exitCode := intTests(ctx, c)
	fmt.Println("ExitCode:", exitCode)

	cancel()
	cc1.Close()
	os.Exit(exitCode)
}
