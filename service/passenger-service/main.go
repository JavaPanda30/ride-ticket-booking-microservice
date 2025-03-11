package main

import (
	"fmt"
	"log"
	"net"

	pb "example.com/uber/Go-microservices/proto/passenger"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Passenger Service...")

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	passengerService := NewPassengerServiceServer()
	pb.RegisterPassengerServiceServer(server, passengerService)

	fmt.Println("Passenger Service running on port 50052...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
