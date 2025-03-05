Plans for 3 Services

Drivers - D
Rides - R
Passengers - P


- P calls function with location and id 
- then it calls R to create ride request
- R receives request and asks D for available drivers
- D gets list 
- R updates a ride status to assigned
- D updates the driver status
- R notifies P to update status

- on pickup D sends update request to R
- on completing R updates ride status and sends D a request to free a driver


- File Structure
```
Go-microservices
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
   │  ├─ driver-service.go
   │  └─ main.go
   ├─ passenger-service
   │  ├─ main.go
   │  └─ passenger-service.go
   └─ ride-service
      ├─ main.go
      └─ ride-service.go

```