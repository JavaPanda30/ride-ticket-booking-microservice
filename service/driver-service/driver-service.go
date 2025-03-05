package main

import (
	"context"
	"log"

	pb "example.com/uber/Go-microservices/proto/driver"
)

type DriverServiceServer struct {
	pb.UnimplementedDriverServiceServer
}

func (s *DriverServiceServer) GetAvailableDrivers(ctx context.Context, req *pb.DriverRequest) (*pb.DriverList, error) {
	log.Printf("Finding available drivers near: %s", req.Location)

	drivers := []*pb.Driver{
		{Id: "D1", Name: "John Doe", Status: "Available"},
		{Id: "D2", Name: "Alice Smith", Status: "Available"},
	}

	return &pb.DriverList{Drivers: drivers}, nil
}

func (s *DriverServiceServer) UpdateDriverStatus(ctx context.Context, req *pb.DriverStatusUpdate) (*pb.DriverResponse, error) {
	log.Printf("Updating driver %s status to %s", req.DriverId, req.Status)
	return &pb.DriverResponse{Success: true}, nil
}
