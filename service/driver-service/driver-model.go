package main 

import (
	"database/sql"
	"fmt"
	"log"
)

type Driver struct {
	DriverID string
	Status   string
}

type DriverModel struct {
	DB *sql.DB
}

func NewDriverModel(db *sql.DB) *DriverModel {
	return &DriverModel{DB: db}
}

func (dm *DriverModel) GetAvailableDrivers() ([]Driver, error) {
	query := `SELECT driver_id, status FROM drivers WHERE status = 'available'`
	rows, err := dm.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching available drivers: %v", err)
	}
	defer rows.Close()

	var drivers []Driver
	for rows.Next() {
		var d Driver
		if err := rows.Scan(&d.DriverID, &d.Status); err != nil {
			log.Printf("error scanning driver row: %v", err)
			continue
		}
		drivers = append(drivers, d)
	}
	return drivers, nil
}

// UpdateDriverStatus updates the status of a driver
func (dm *DriverModel) UpdateDriverStatus(driverID, status string) error {
	query := `UPDATE drivers SET status = $1 WHERE driver_id = $2`
	_, err := dm.DB.Exec(query, status, driverID)
	if err != nil {
		return fmt.Errorf("error updating driver status: %v", err)
	}
	return nil
}
