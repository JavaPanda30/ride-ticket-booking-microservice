package main

import (
	"fmt"
	"log"
	"net"

	pb "example.com/uber/Go-microservices/proto/driver"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDriverServiceServer(grpcServer, &DriverServiceServer{})

	fmt.Println("Driver Service is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
