package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"

	pb "github.com/brisk84/gophkeeper/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

type server struct {
	pb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	lis, err := net.Listen("tcp", ":4343")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	// grpcServer := grpc.NewServer()

	pb.RegisterGreeterServer(grpcServer, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:4343")
	log.Fatal(grpcServer.Serve(lis))
	// grpcServer := grpc.NewServer(
	// 	grpc.Creds(tlsCredentials),
	// 	grpc.UnaryInterceptor(interceptor.Unary()),
	// 	grpc.StreamInterceptor(interceptor.Stream()),
	// )
}
