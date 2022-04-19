package models

import "time"

type Ride struct {
	idRide   int64
	created  time.Time
	finished time.Time

	user    User
	vehicle Vehicle
}

func (ride *Ride) Constructor(user User, vehicle Vehicle) {
	ride.created = time.Now()
	ride.user = user
	ride.vehicle = vehicle
}

func (ride *Ride) GetIdRide() int64 {
	return ride.idRide
}

func (ride *Ride) SetIdUser(idRide int64) {
	ride.idRide = idRide
}

func (ride *Ride) GetCreated() time.Time {
	return ride.created
}

func (ride *Ride) SetCreated(created time.Time) {
	ride.created = created
}

func (ride *Ride) GetFinished() time.Time {
	return ride.finished
}

func (ride *Ride) SetFinished(finished time.Time) {
	ride.finished = finished
}

func (ride *Ride) GetUser() User {
	return ride.user
}

func (ride *Ride) SetUser(user User) {
	ride.user = user
}

func (ride *Ride) GetVehicle() Vehicle {
	return ride.vehicle
}

func (ride *Ride) SetVehicle(vehicle Vehicle) {
	ride.vehicle = vehicle
}

func (ride *Ride) CheckRideNotFinished() bool {
	return ride.finished.IsZero()
}
