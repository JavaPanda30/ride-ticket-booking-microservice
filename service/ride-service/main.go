package main

import (
	"fmt"
	"log"
	"net"

	pb "example.com/uber/Go-microservices/proto/ride"
	"example.com/uber/db"
	"google.golang.org/grpc"
)

func main() {
	// Initialize DB
	db.InitDB()

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRideServiceServer(grpcServer, &RideServiceServer{})

	fmt.Println("Ride Service is running on port 50053...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
