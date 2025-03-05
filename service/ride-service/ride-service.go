package main

import (
	"context"
	"log"

	pb "example.com/uber/Go-microservices/proto/ride"
)

type RideServiceServer struct {
	pb.UnimplementedRideServiceServer
}

func (s *RideServiceServer) CreateRide(ctx context.Context, req *pb.RideRequest) (*pb.RideResponse, error) {
	log.Printf("Creating ride for passenger %s at location %s", req.PassengerId, req.PickupLocation)
	return &pb.RideResponse{RideId: "R12345", Success: true}, nil
}

func (s *RideServiceServer) UpdateRideStatus(ctx context.Context, req *pb.RideStatusUpdate) (*pb.RideResponse, error) {
	log.Printf("Updating ride %s status to %s", req.RideId, req.Status)
	return &pb.RideResponse{RideId: req.RideId, Success: true}, nil
}
