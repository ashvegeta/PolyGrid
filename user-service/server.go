package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ashvegeta/user-service/generated"
	"google.golang.org/grpc"
)

type ConnOps struct {
	Network string
	Addr    string
}

type UserServer struct {
	pb.UnimplementedUserServer
	ConnOps
	grpcServer  *grpc.Server
	tcpListener net.Listener
}

func (s *UserServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

// Start the server
func (s *UserServer) Start() {
	// Check if Network and Addr are set
	if s.Network == "" || s.Addr == "" {
		log.Fatal("Network and Addr must be set")
	}

	// Listen on the network address
	lis, err := net.Listen(s.Network, s.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Server listening on %s\n", lis.Addr().String())

	// Create a new gRPC server and listen on the network address
	// for incoming requests
	server := grpc.NewServer()
	pb.RegisterUserServer(server, s)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *UserServer) Close() {
	// Close the gRPC server
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
		fmt.Println("gRPC server gracefully stopped")
	}

	// Close the TCP listener
	if s.tcpListener != nil {
		if err := s.tcpListener.Close(); err != nil {
			log.Fatalf("failed to close TCP listener: %v", err)
		}
		fmt.Println("TCP listener closed")
	}
}
