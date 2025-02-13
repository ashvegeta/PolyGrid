# Microservices Project

## About

This project is a proof of concept (PoC) for a microservices-based architecture that uses Kafka for messaging and gRPC for inter-service communication. The project consists of three main services: a producer service, a consumer service, and an analytics service.

## Tech Stack

- **Programming Languages**: Go, Python, JavaScript (Node.js)
- **Messaging**: Apache Kafka
- **gRPC**: For inter-service communication
- **Docker**: For containerization
- **Environment Variables**: Managed using `.env` files

## Services

### Producer Service

- **Language**: Node.js
- **Description**: Produces messages to a Kafka topic and logs the messages using a gRPC call to the analytics service.
- **Dependencies**: `kafkajs`, `@grpc/grpc-js`, `@grpc/proto-loader`, `dotenv`
- **File**: `producer-service/producer.js`

### Consumer Service

- **Language**: Go
- **Description**: Consumes messages from a Kafka topic and logs the messages using a gRPC call to the analytics service.
- **Dependencies**: `sarama`, `google.golang.org/grpc`
- **File**: `consumer-service/main.go`

### Analytics Service

- **Language**: Python
- **Description**: Receives log messages via gRPC and stores them in log files.
- **Dependencies**: `grpcio`, `grpcio-tools`, `python-dotenv`
- **File**: `analytics-service/server.py`

## Setup

- ### Prerequisites

  - Docker
  - Node.js
  - Go
  - Python

- ### Environment Variables

  - Create a `.env` file in the root directory with the following content:
  - add
    ```
    KAFKA_ADDR=localhost:9092
    ANALYTICS_GRPC_ADDR=localhost:8080
    ```

- ### Generate gRPC Code

  - For the analytics service (Python):

    ```sh
    cd analytics-service
    make gen
    ```

  - For the consumer service (Go):

    ```sh
    cd consumer-service
    make gen
    ```

  - For the producer service (Node.js):

    ```sh
    cd producer-service
    make gen
    ```

- ### Running the Services

  - Setup and start Kafka broker:

    if first time setting up (run in root folder):

    ```sh
    make setup-kafka
    ```

    else, run:

    ```sh
    make run
    ```

  - Start the analytics service:

    ```sh
    cd analytics-service
    make run
    ```

  - Start the consumer service:

    ```sh
    cd consumer-service
    make run
    ```

  - Start the producer service::

    ```sh
    cd producer-service
    make run
    ```

## Architecture

The project follows a microservices architecture with the following components:

1. **Producer Service**: Sends messages to a Kafka topic and logs them using gRPC.

2. **Consumer Service**: Consumes messages from the Kafka topic and logs them using gRPC.

3. **Analytics Service**: Receives log messages via gRPC and stores them in log files.
