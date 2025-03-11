package main

import (
	"context"

	pb "example.com/uber/Go-microservices/proto/ride"
)

type RideServiceServer struct {
	pb.UnimplementedRideServiceServer
}

func (s *RideServiceServer) CreateRide(ctx context.Context, req *pb.RideRequest) (*pb.RideResponse, error) {
	ride, err := CreateRide(req.PassengerId, req.PickupLocation, req.DropLocation)
	if err != nil {
		return nil, err
	}
	driverID := "driver_123"
	if err := UpdateRideStatus(ride.RideID, "assigned"); err != nil {
		return nil, err
	}

	return &pb.RideResponse{
		RideId:   ride.RideID,
		Status:   "assigned",
		DriverId: driverID,
		Success:  true,
	}, nil
}

func (s *RideServiceServer) UpdateRideStatus(ctx context.Context, req *pb.RideStatusUpdate) (*pb.RideResponse, error) {
	err := UpdateRideStatus(req.RideId, req.Status)
	if err != nil {
		return nil, err
	}

	return &pb.RideResponse{
		RideId:  req.RideId,
		Status:  req.Status,
		Success: true,
	}, nil
}
