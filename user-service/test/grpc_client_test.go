package test

import (
	"context"
	"log"
	"testing"

	pb "github.com/ashvegeta/user-service/generated"
	srv "github.com/ashvegeta/user-service/server"
	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {
	//Create a new server
	a := srv.UserServer{}
	a.ConnOps = srv.ConnOps{
		Network: "tcp",
		Addr:    ":8080",
	}
	go a.Start()
	defer a.Close()

	// Create a new client
	conn, err := grpc.NewClient("localhost:8080", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	// Call the remote method
	req := &pb.HelloRequest{Name: "Ash"}
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		t.Fatalf("failed to call: %v", err)
	}
	log.Printf("Response: %s", resp.Message)

	// Add assertions to verify the response
	expectedMessage := "Hello Ash"
	if resp.Message != expectedMessage {
		t.Errorf("expected %s, got %s", expectedMessage, resp.Message)
	}
}
