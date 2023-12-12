package main

import (
	"context"
	"flag"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"database/sql"
	"log"
	"net"

	"github.com/sikemausa/micro-service-example/internal/handler"
	"github.com/sikemausa/micro-service-example/internal/repository/postgres"
	"github.com/sikemausa/micro-service-example/internal/service"
	user_v1 "github.com/sikemausa/micro-service-example/pkg/proto/user/v1"
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

	user_v1.RegisterUserServiceServer(grpcServer, handler.NewUserServiceServer(userService))

	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	flag.Parse()
	if err := startHTTPServer(); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}

func startHTTPServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := user_v1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}
