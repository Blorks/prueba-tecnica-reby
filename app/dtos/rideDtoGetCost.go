package dtos

import (
	"reby/app/models"
)

type RideDtoGetCost struct {
	RideDtoGet
	Cost int `json:"cost"`
}

func (rideDTO *RideDtoGetCost) Constructor(ride models.Ride, cost int) {
	rideDTO.IdRide = ride.IdRide
	rideDTO.Created = ride.Created
	rideDTO.Finished = ride.Finished

	rideDTO.IdUser = ride.IdUser
	rideDTO.IdVehicle = ride.IdVehicle

	rideDTO.Cost = cost
}
