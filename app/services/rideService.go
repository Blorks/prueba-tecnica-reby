package services

import (
	"math"
	"reby/app/dtos"
	"reby/app/models"
	"reby/app/repositories"
	"time"
)

const unlock_price int = 100
const minute_price int = 18

func calculateCost(ride models.Ride) int {
	if ride.Finished.IsZero() {
		panic("It is not possible to calculate the cost of a ride that has not yet finished")
	}

	diff := ride.Finished.Sub(ride.Created)
	minutes := int(math.Ceil(diff.Minutes()))

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

func FinishRide(idRide int) dtos.RideDtoGetCost {
	ride := repositories.GetRide(idRide)

	if !ride.CheckRideNotFinished() {
		panic("The ride that is trying to end is already over")
	}

	ride.Finished = time.Now()

	ride = repositories.SaveRide(ride)

	response := dtos.RideDtoGetCost{}
	response.Constructor(ride, calculateCost(ride))

	return response
}
