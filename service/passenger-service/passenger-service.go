package main

import (
	"context"
	"database/sql"
	"fmt"

	pb "example.com/uber/Go-microservices/proto/passenger"
	"example.com/uber/db"
)

type PassengerServiceServer struct {
	pb.UnimplementedPassengerServiceServer
	DB *sql.DB
}

func NewPassengerServiceServer() *PassengerServiceServer {
	database, err := db.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
	return &PassengerServiceServer{DB: database}
}

func (s *PassengerServiceServer) RequestRide(ctx context.Context, req *pb.PassengerRequest) (*pb.PassengerResponse, error) {
	query := `INSERT INTO passengers (passenger_id, pickup_location, drop_location) VALUES ($1, $2, $3)`
	_, err := s.DB.Exec(query, req.PassengerId, req.PickupLocation, req.DropLocation)
	if err != nil {
		return &pb.PassengerResponse{Success: false}, fmt.Errorf("failed to create passenger: %v", err)
	}

	rideID, err := s.CreateRideRequest(req)
	if err != nil {
		return &pb.PassengerResponse{Success: false}, err
	}

	return &pb.PassengerResponse{RideId: rideID, Success: true}, nil
}

func (s *PassengerServiceServer) CreateRideRequest(req *pb.PassengerRequest) (string, error) {
	rideID := "ride_" + req.PassengerId
	return rideID, nil
}

func (s *PassengerServiceServer) UpdatePassengerStatus(ctx context.Context, req *pb.PassengerStatusUpdate) (*pb.PassengerResponse, error) {
	fmt.Printf("Updating passenger %s status to %s\n", req.PassengerId, req.Status)
	return &pb.PassengerResponse{Success: true}, nil
}
