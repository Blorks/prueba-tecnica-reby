package services

import (
	"math"
	"reby/app/dtos"
	"reby/app/models"
	"reby/app/repositories"
	"time"
)

const unlock_price int64 = 100
const minute_price int64 = 18

func calculateCost(ride models.Ride) int64 {
	if ride.GetFinished().IsZero() {
		panic("It is not possible to calculate the cost of a ride that has not yet finished")
	}

	diff := ride.GetFinished().Sub(ride.GetCreated())
	minutes := int64(math.Ceil(diff.Minutes()))

	return unlock_price + minutes*minute_price
}

func InitRide(rideDto dtos.RideDtoPost) dtos.RideDtoGet {
	user := repositories.GetUser(rideDto.IdUser)
	vehicle := repositories.GetVehicle(rideDto.IdVehicle)

	if !user.CheckUserBalance(unlock_price) {
		panic("The user's balance is too low to start the ride")
	}

	if !vehicle.CheckVehicleFree() {
		panic("The used vehicle is not available at this time")
	}

	ride := models.Ride{}
	ride.Constructor(user, vehicle)

	ride = repositories.SaveRide(ride)

	response := dtos.RideDtoGet{}
	response.Constructor(ride)

	return response
}

func FinishRide(idRide int64) dtos.RideDtoGetCost {
	ride := repositories.GetRide(idRide)

	if !ride.CheckRideNotFinished() {
		panic("The ride that is trying to end is already over")
	}

	ride.SetFinished(time.Now())

	repositories.SaveRide(ride)

	response := dtos.RideDtoGetCost{}
	response.Constructor(ride, calculateCost(ride))

	return response
}
