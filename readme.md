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
