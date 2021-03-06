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

type RideService struct {
	userRepository    repositories.UserRepository
	vehicleRepository repositories.VehicleRepository
	rideRepository    repositories.RideRepository
}

func NewRideService(userRepository repositories.UserRepository, vehicleRepository repositories.VehicleRepository, rideRepository repositories.RideRepository) *RideService {
	return &RideService{userRepository: userRepository, vehicleRepository: vehicleRepository, rideRepository: rideRepository}
}

/**
Calculate the cost of a ride based on the following formula:

cost = unlock_price + minutes*minute_price

- If the ride for which you want to calculate the cost has not finished yet, return an error
*/
func calculateCost(ride models.Ride) int {
	if ride.Finished.IsZero() {
		panic("It is not possible to calculate the cost of a ride that has not yet finished")
	}

	diff := ride.Finished.Sub(ride.Created)
	minutes := int(math.Ceil(diff.Minutes()))

	return unlock_price + minutes*minute_price
}

/**
Create a ride

- If the user does not have enough balance to start the ride, return an error
- If the vehicle you want to use is already in use, it returns an error

Returns an object with the information of the inserted Ride
*/
func (s *RideService) InitRide(rideDto dtos.RideDtoPost) dtos.RideDtoGet {
	user, _ := s.userRepository.GetUser(rideDto.IdUser)
	vehicle, _ := s.vehicleRepository.GetVehicle(rideDto.IdVehicle)

	if !user.CheckUserBalance(unlock_price) {
		panic("The user's balance is too low to start the ride")
	}

	if !vehicle.CheckVehicleFree() {
		panic("The used vehicle is not available at this time")
	}

	ride := models.Ride{}
	ride.Constructor(user, vehicle)

	ride = s.rideRepository.SaveRide(ride)

	response := dtos.RideDtoGet{}
	response.Constructor(ride)

	return response
}

/**
Finish a ride

- If the Ride you want to finish is already finished, it returns an error

Returns an object with the information of the finished Ride and his cost
*/
func (s *RideService) FinishRide(idRide int) dtos.RideDtoGetCost {
	ride, _ := s.rideRepository.GetRide(idRide)

	if !ride.CheckRideNotFinished() {
		panic("The ride that is trying to end is already over")
	}

	ride.Finished = time.Now()

	ride = s.rideRepository.SaveRide(ride)

	response := dtos.RideDtoGetCost{}
	response.Constructor(ride, calculateCost(ride))

	return response
}
