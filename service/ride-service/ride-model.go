package main

import (
	"database/sql"
	"fmt"

	"example.com/uber/db"
)

type Ride struct {
	RideID         string
	PassengerID    string
	PickupLocation string
	DropLocation   string
	DriverID       sql.NullString
	Status         string
}

func CreateRide(passengerID, pickupLocation, dropLocation string) (*Ride, error) {
	query := `INSERT INTO rides (passenger_id, pickup_location, status) 
	          VALUES ($1, $2, 'requested') RETURNING ride_id;`
	var rideID string
	err := db.DB.QueryRow(query, passengerID, pickupLocation).Scan(&rideID)
	if err != nil {
		return nil, fmt.Errorf("failed to create ride: %v", err)
	}

	return &Ride{
		RideID:         rideID,
		PassengerID:    passengerID,
		PickupLocation: pickupLocation,
		DropLocation:   dropLocation,
		Status:         "requested",
	}, nil
}

func UpdateRideStatus(rideID, status string) error {
	query := `UPDATE rides SET status = $1 WHERE ride_id = $2;`
	_, err := db.DB.Exec(query, status, rideID)
	if err != nil {
		return fmt.Errorf("failed to update ride status: %v", err)
	}
	return nil
}
