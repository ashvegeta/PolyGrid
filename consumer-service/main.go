package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
)

func consume(wg *sync.WaitGroup) {
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

consumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message: %s\n", string(msg.Value))
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

	// Add a task to the WaitGroup
	wg.Add(1)

	// Start the consumer in a separate goroutine
	go consume(&wg)

	// Wait for the consumer to complete
	wg.Wait()
}
