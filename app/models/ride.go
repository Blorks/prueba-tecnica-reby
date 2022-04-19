package models

import "time"

type Ride struct {
	IdRide   int `gorm:"primaryKey"`
	Created  time.Time
	Finished time.Time

	IdUser int
	User   User `gorm:"foreignKey:IdUser"`

	IdVehicle int
	Vehicle   Vehicle `gorm:"foreignKey:IdVehicle"`
}

func (ride *Ride) Constructor(user User, vehicle Vehicle) {
	ride.Created = time.Now()
	ride.Finished = ride.Created

	ride.AppendAssociations(user, vehicle)
}

func (ride *Ride) AppendAssociations(user User, vehicle Vehicle) {
	ride.IdUser = user.IdUser
	ride.User = user

	ride.IdVehicle = vehicle.IdVehicle
	ride.Vehicle = vehicle
}

func (ride *Ride) CheckRideNotFinished() bool {
	return ride.Finished.Equal(ride.Created)
}
