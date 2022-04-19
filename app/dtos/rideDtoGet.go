package dtos

import (
	"reby/app/models"
	"time"
)

type RideDtoGet struct {
	IdRide   int64     `json:"idRide"`
	Created  time.Time `json:"created"`
	Finished time.Time `json:"finished"`
	RideDto
}

func (rideDTO *RideDtoGet) Constructor(ride models.Ride) {
	rideDTO.IdRide = ride.GetIdRide()
	rideDTO.Created = ride.GetCreated()
	rideDTO.Finished = ride.GetFinished()

	user := ride.GetUser()
	rideDTO.IdUser = user.GetIdUser()

	vehicle := ride.GetVehicle()
	rideDTO.IdVehicle = vehicle.GetIdVehicle()
}
