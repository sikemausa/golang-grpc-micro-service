# Golang gRPC Micro-Service

This repository hosts a micro-service developed in Go, utilizing gRPC for efficient and scalable inter-service communication. It's designed to showcase the implementation of a gRPC server and client in Go.

## Features

- gRPC Server: Implements a gRPC server for handling remote procedure calls.
- gRPC Client: A client that communicates with the server using gRPC.
- Go Implementation: The service is written in Go, offering high performance and concurrency support.

## Getting Started

These instructions will guide you through getting a copy of the project up and running on your local machine for development and testing purposes.
Prerequisites

- Go (Version specified in go.mod)
- gRPC
- Protocol Buffers (protoc)
- kind
- kubectl
- Docker Desktop

## Installation

Clone the Repository
    
    git clone https://github.com/sikemausa/golang-grpc-micro-service.git

Navigate to the Project Directory

    cd golang-grpc-micro-service

Create a local kubernetes cluster with kind

    kind create cluster

Create the postgres database

    kubectl apply -f k8s/pg.yaml

Get the DB password

    kubectl get secret pg-db-app -o go-template='{{.data.password | base64decode}}'

Set the DB URL

    export DB_URL="postgresql://app:<<YOUR_PASSWORD>>@localhost:7000/app?sslmode=disable"

Run the DB migrations

    make migrate_up

Port forward from localhost to the database on the kubernetes cluster

    kubectl port-forward pg-db-1 7000:5432

Updating/Installing Dependencies

    make update_dependencies

Building the Service

    make build

Running the Server

    make run

## Interacting with the Service

This section demonstrates how to make requests to the server using both HTTP and gRPC.

### Making HTTP Requests

The service exposes HTTP endpoints that can be interacted with using standard HTTP methods. Below are examples of how to interact with these endpoints using curl.
List Users

    Endpoint: GET /v1/users

    curl -X GET http://localhost:8080/v1/users

Create User

    Endpoint: POST /v1/users

    curl -X POST http://localhost:8080/v1/users \
         -H "Content-Type: application/json" \
         -d '{"email": "example@email.com", "name": "John Doe"}'

Get User

    Endpoint: GET /v1/users/{id}

    curl -X GET http://localhost:8080/v1/users/{id}

Delete User

    Endpoint: DELETE /v1/users/{id}

    curl -X DELETE http://localhost:8080/v1/users/{id}

Update User

    Endpoint: PATCH /v1/users/{id}

    curl -X PATCH http://localhost:8080/v1/users/{id} \
         -H "Content-Type: application/json" \
         -d '{"email": "newemail@email.com", "name": "Jane Doe"}'

### Making gRPC Requests

To interact with the service using gRPC, you'll need a gRPC client that is compatible with the service's protocol. Below is a general outline of how to make gRPC requests to the service.

#### Example gRPC Request

Here's a pseudocode example of how a gRPC client might interact with the service:

    import (
        "context"
        "log"
    
        "google.golang.org/grpc"
        pb "github.com/sikemausa/micro-service-example/pb/v1"
    )

    func main() {
        conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
        if err != nil {
            log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
    
        c := pb.NewUserServiceClient(conn)
    
        // Example: List Users
        users, err := c.ListUsers(context.Background(), &pb.ListUsersRequest{})
        if err != nil {
            log.Fatalf("could not list users: %v", err)
        }
        log.Printf("Users: %v", users)
    }
