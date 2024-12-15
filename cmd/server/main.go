package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "github.com/irtza33/user_service/internal/config"
    "github.com/irtza33/user_service/internal/service"
    "github.com/irtza33/user_service/pkg/logger"
    "github.com/irtza33/user_service/pkg/database"
    "github.com/irtza33/user_service/proto/user"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Initialize logger
    log := logger.New(cfg.LogLevel)

    // Connect to the database
    db, err := database.Connect(cfg.Database)
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    defer db.Close()

    // Create a new gRPC server
    grpcServer := grpc.NewServer()

    // Register user service
    userService := service.NewUserService(db)
    user.RegisterUserServiceServer(grpcServer, userService)

    // Start listening for incoming connections
    listener, err := net.Listen("tcp", cfg.Server.Address)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    log.Printf("Starting gRPC server on %s", cfg.Server.Address)
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}