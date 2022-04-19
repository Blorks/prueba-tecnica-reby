package dtos

import (
	"reby/app/models"
	"time"
)

type RideDtoGet struct {
	IdRide   int       `json:"idRide"`
	Created  time.Time `json:"created"`
	Finished time.Time `json:"finished"`
	RideDto
}

func (rideDTO *RideDtoGet) Constructor(ride models.Ride) {
	rideDTO.IdRide = ride.IdRide
	rideDTO.Created = ride.Created
	rideDTO.Finished = ride.Finished

	rideDTO.IdUser = ride.IdUser
	rideDTO.IdVehicle = ride.IdVehicle
}
