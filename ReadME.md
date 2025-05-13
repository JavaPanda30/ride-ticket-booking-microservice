# Ride Ticket Booking Microservice

## Overview

This project implements a microservice-based system for booking ride tickets. It consists of three services:

- **Driver Service**: Manages driver availability and status.
- **Passenger Service**: Handles passenger ride requests.
- **Ride Service**: Coordinates ride creation and status updates.

## Prerequisites

1. Install [Go](https://golang.org/dl/).
2. Install [PostgreSQL](https://www.postgresql.org/download/).
3. Install the required Go modules:

   
bash
   go mod tidy


## Setting Up PostgreSQL

1. Create three databases:

   - ride_service
   - driver_service
   - passenger_service

2. Update the connection strings in the service files if necessary.

3. Each service will automatically create the required tables on startup.

## Running the Services

1. Start the **Driver Service**:

   
bash
   go run service/driver-service/main.go


2. Start the **Passenger Service**:

   
bash
   go run service/passenger-service/main.go


3. Start the **Ride Service**:

   
bash
   go run service/ride-service/main.go


## Workflow

1. A passenger requests a ride through the Passenger Service.
2. The Ride Service validates the passenger and assigns an available driver.
3. The Driver Service updates the driver's status.
4. The Ride Service tracks the ride's status and notifies the Passenger Service of updates.

## File Structure

plaintext
ride-ticket-booking-microservice
├─ frontend
│  └─ index.html
├─ Go-microservices
│  └─ proto
│     ├─ driver
│     │  ├─ driver.pb.go
│     │  └─ driver_grpc.pb.go
│     ├─ passenger
│     │  ├─ passenger.pb.go
│     │  └─ passenger_grpc.pb.go
│     └─ ride
│        ├─ ride.pb.go
│        └─ ride_grpc.pb.go
├─ go.mod
├─ go.sum
├─ proto
│  ├─ driver.proto
│  ├─ passenger.proto
│  └─ ride.proto
├─ readme.md
└─ service
   ├─ driver-service
   │  ├─ Dockerfile
   │  ├─ driver-service.go
   │  └─ main.go
   ├─ passenger-service
   │  ├─ Dockerfile
   │  ├─ main.go
   │  └─ passenger-service.go
   └─ ride-service
      ├─ Dockerfile
      ├─ go.sum
      ├─ main.go
      └─ ride-service.go
      