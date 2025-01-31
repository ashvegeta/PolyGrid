package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"google.golang.org/grpc"

	"github.com/IBM/sarama"
	pb "github.com/ashvegeta/PolyGrid/generated"
)

func consume(wg *sync.WaitGroup, gRPCClient pb.AnalyticsServiceClient) {
	defer wg.Done() // Notify main that this goroutine is done

	// Configure Kafka consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Define Kafka brokers
	brokers := []string{"localhost:9092"} // Replace with your Kafka broker's address

	// Create a new consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Printf("Error closing consumer: %s\n", err)
		}
	}()

	// Subscribe to Kafka topic
	topic := "analytics"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			fmt.Printf("Error closing partition consumer: %s\n", err)
		}
	}()

	// Handle messages
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// main channel switch
consumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			{
				fmt.Printf("Received message from kafka broker: %s\n", string(msg.Value))
				// Call the remote gRPC method
				req := &pb.SendLogRequest{Message: string(msg.Value), SenderType: "consumer"}

				resp, err := gRPCClient.SendLog(context.Background(), req)
				if err != nil {
					fmt.Printf("ERROR: failed to call gRPC server: %v\n", err)
					fmt.Println("Continuing consumption...")
					continue
				}
				fmt.Printf("Received gRPC ACK: %s\n\n", resp.Message)
			}
		case err := <-partitionConsumer.Errors():
			fmt.Printf("Error: %s\n", err)
		case <-signals:
			fmt.Println("Interrupt signal received. Shutting down...")
			break consumerLoop
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// open new connection to gRPC server
	conn, err := grpc.NewClient("localhost:8080", grpc.WithInsecure())
	gRPCClient := pb.NewAnalyticsServiceClient(conn)
	if err != nil {
		fmt.Printf("ERROR : failed to dial gRPC server: %v\n", err)
	}
	defer conn.Close()

	// Add a task to the WaitGroup
	wg.Add(1)

	// Start the consumer in a separate goroutine
	go consume(&wg, gRPCClient)

	// Wait for the consumer to complete
	wg.Wait()
}
