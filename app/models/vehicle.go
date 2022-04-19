package models

import (
	"reby/app/models/enums"
)

type Vehicle struct {
	IdVehicle int `gorm:"primaryKey"`
	Name      string
	State     enums.VehicleState
}

func (vehicle *Vehicle) Constructor(name string, state enums.VehicleState) {
	vehicle.Name = name
	vehicle.State = state
}

func (vehicle *Vehicle) CheckVehicleFree() bool {
	return vehicle.State == enums.Free
}
