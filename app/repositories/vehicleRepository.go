package repositories

import (
	"reby/app/models"
)

func GetVehicle(idVehicle int) models.Vehicle {
	vehicle := models.Vehicle{}

	Database.First(&vehicle, idVehicle)

	return vehicle
}
