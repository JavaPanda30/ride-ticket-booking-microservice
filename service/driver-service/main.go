package main

import (
	"fmt"
	"log"
	"net"

	"example.com/uber/db"
	pb "example.com/uber/Go-microservices/proto/driver"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Driver Service...")

	DB, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer DB.Close()

	driverModel := NewDriverModel(DB)

	grpcServer := grpc.NewServer()
	driverService := NewDriverService(driverModel)
	pb.RegisterDriverServiceServer(grpcServer, driverService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Driver Service is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
