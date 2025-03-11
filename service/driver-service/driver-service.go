package main

import (
	"context"
	"fmt"

	pb "example.com/uber/Go-microservices/proto/driver"
)

type DriverServiceServer struct {
	pb.UnimplementedDriverServiceServer
	DriverModel *DriverModel
}

func NewDriverService(dm *DriverModel) *DriverServiceServer {
	return &DriverServiceServer{DriverModel: dm}
}

func (s *DriverServiceServer) GetAvailableDrivers(ctx context.Context, req *pb.DriverRequest) (*pb.DriverResponse, error) {
	drivers, err := s.DriverModel.GetAvailableDrivers()
	if err != nil {
		return nil, err
	}

	var driverList []*pb.Driver
	for _, d := range drivers {
		driverList = append(driverList, &pb.Driver{
			DriverId: d.DriverID,
			Status:   d.Status,
		})
	}

	return &pb.DriverResponse{Drivers: driverList}, nil
}

func (s *DriverServiceServer) UpdateDriverStatus(ctx context.Context, req *pb.DriverStatusUpdate) (*pb.DriverResponse, error) {
	err := s.DriverModel.UpdateDriverStatus(req.DriverId, req.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to update driver status: %v", err)
	}
	return &pb.DriverResponse{Success: true}, nil
}
