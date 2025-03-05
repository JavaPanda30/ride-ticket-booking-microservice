package main

import (
	"fmt"
	"log"
	"net"

	pb "example.com/uber/Go-microservices/proto/passenger"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen on port 50053: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPassengerServiceServer(grpcServer, &PassengerServiceServer{})

	fmt.Println("Passenger Service is running on port 50053...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
