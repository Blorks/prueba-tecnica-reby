package models

import "reby/app/models/enums"

type Vehicle struct {
	idVehicle int64
	name      string
	state     enums.VehicleState
}

func (vehicle *Vehicle) Constructor(name string, state enums.VehicleState) {
	vehicle.name = name
	vehicle.state = state
}

func (vehicle *Vehicle) GetIdVehicle() int64 {
	return vehicle.idVehicle
}

func (vehicle *Vehicle) SetIdVehicle(idVehicle int64) {
	vehicle.idVehicle = idVehicle
}

func (vehicle *Vehicle) GetName() string {
	return vehicle.name
}

func (vehicle *Vehicle) SetName(name string) {
	vehicle.name = name
}

func (vehicle *Vehicle) GetState() enums.VehicleState {
	return vehicle.state
}

func (vehicle *Vehicle) SetState(state enums.VehicleState) {
	vehicle.state = state
}

func (vehicle *Vehicle) CheckVehicleFree() bool {
	return vehicle.state == enums.Free
}
