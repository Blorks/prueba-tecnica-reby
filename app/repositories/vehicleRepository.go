package repositories

import (
	"reby/app/models"
)

func GetVehicle(idVehicle int64) models.Vehicle {
	vehicle := models.Vehicle{}

	Database.First(&vehicle, idVehicle)

	return vehicle
}
