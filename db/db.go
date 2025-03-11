package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	connStr := "user=aman password=admin port=1234 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Check if the connection is alive
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database is unreachable: %v", err)
	}

	fmt.Println("Connected to the database successfully!")
	return db, nil
}

func createPassengersTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS passengers (
		passenger_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		pickup_location TEXT NOT NULL,
		drop_location TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	return err
}

func createDriversTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS drivers (
		driver_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		status TEXT NOT NULL,
		passenger_assigned UUID UNIQUE,
		FOREIGN KEY(passenger_assigned) REFERENCES passengers(passenger_id) ON DELETE SET NULL
	);`
	_, err := db.Exec(query)
	return err
}

func createRidesTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS rides (
		ride_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		pickup_location TEXT NOT NULL,
		driver_assigned UUID UNIQUE,
		passenger_id UUID,
		status TEXT NOT NULL,
		FOREIGN KEY(driver_assigned) REFERENCES drivers(driver_id) ON DELETE CASCADE,
		FOREIGN KEY(passenger_id) REFERENCES passengers(passenger_id) ON DELETE CASCADE
	);`
	_, err := db.Exec(query)
	return err
}

func SetupDatabase(db *sql.DB) error {
	if err := createPassengersTable(db); err != nil {
		return fmt.Errorf("error creating passengers table: %v", err)
	}
	if err := createDriversTable(db); err != nil {
		return fmt.Errorf("error creating drivers table: %v", err)
	}
	if err := createRidesTable(db); err != nil {
		return fmt.Errorf("error creating rides table: %v", err)
	}
	fmt.Println("Database tables created successfully!")
	return nil
}

func InitDB() {
	var err error
	DB, err = ConnectDB()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	if err := SetupDatabase(DB); err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
}
