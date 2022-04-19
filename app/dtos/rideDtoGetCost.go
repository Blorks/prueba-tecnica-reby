package dtos

import (
	"reby/app/models"
)

type RideDtoGetCost struct {
	RideDtoGet
	Cost int64 `json:"cost"`
}

func (rideDTO *RideDtoGetCost) Constructor(ride models.Ride, cost int64) {
	rideDTO.IdRide = ride.GetIdRide()
	rideDTO.Created = ride.GetCreated()
	rideDTO.Finished = ride.GetFinished()

	user := ride.GetUser()
	rideDTO.IdUser = user.GetIdUser()

	vehicle := ride.GetVehicle()
	rideDTO.IdVehicle = vehicle.GetIdVehicle()

	rideDTO.Cost = cost
}
