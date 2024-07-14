package main

import (
	"context"
	"log"
	"net"

	"github.com/marcelofabianov/pejota/bootstrap"
	"github.com/marcelofabianov/pejota/internal/user"
	server "github.com/marcelofabianov/pejota/internal/user/adapter/grpc"
	userpb "github.com/marcelofabianov/pejota/internal/user/adapter/grpc/generated"
	"github.com/marcelofabianov/pejota/internal/user/port"
	"github.com/marcelofabianov/pejota/pkg/postgres"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load config
	if err := bootstrap.Load(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	config := bootstrap.NewConfig()
	ctx := context.Background()

	// Connect to database
	db, err := postgres.ConnectDB(*config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close(ctx)

	// Create user container
	userContainer := user.NewUserContainer(db)

	// Invoke user service application
	var userServiceApp port.UserServiceApplication
	err = userContainer.Invoke(func(service port.UserServiceApplication) {
		userServiceApp = service
	})
	if err != nil {
		log.Fatalf("Error invoking user service: %v", err)
	}

	// gRPC server
	grpcServer := grpc.NewServer()
	grpcAddress := config.GrpcAddress

	// Register user service server
	userServiceServer := server.NewUserServiceServer(userServiceApp)
	userpb.RegisterUserServiceServer(grpcServer, userServiceServer)

	// Reflection
	reflection.Register(grpcServer)

	// Create listener
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start gRPC server
	log.Printf("gRPC server listening on %s", grpcAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
