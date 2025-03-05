package main

import (
	"context"
	"log"

	pb "example.com/uber/Go-microservices/proto/passenger"
)

type PassengerServiceServer struct {
	pb.UnimplementedPassengerServiceServer
}

func (s *PassengerServiceServer) RequestRide(ctx context.Context, req *pb.PassengerRideRequest) (*pb.RideResponse, error) {
	log.Printf("Passenger %s requested a ride from %s", req.PassengerId, req.Location)
	return &pb.RideResponse{RideId: "R98765", Success: true}, nil
}
