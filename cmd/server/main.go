package main

import (
	"os"

	_ "github.com/lib/pq"
	"google.golang.org/grpc/reflection"

	"database/sql"
	"log"
	"net"

	"github.com/sikemausa/micro-service-example/internal/handler"
	"github.com/sikemausa/micro-service-example/internal/repository/postgres"
	"github.com/sikemausa/micro-service-example/internal/service"
	"github.com/sikemausa/micro-service-example/pkg/proto/v1"
	"google.golang.org/grpc"
)

func main() {
	dbConnectionString := os.Getenv("DATABASE_URL")
	if dbConnectionString == "" {
		log.Fatalf("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open(
		"postgres",
		dbConnectionString,
	)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	userRepo := postgres.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterUserServiceServer(grpcServer, handler.NewUserServiceServer(userService))

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
