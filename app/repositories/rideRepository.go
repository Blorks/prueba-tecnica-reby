package repositories

import (
	"reby/app/models"
)

func GetRide(idRide int64) models.Ride {
	ride := models.Ride{}

	Database.First(&ride, idRide)

	return ride
}

func SaveRide(ride models.Ride) models.Ride {
	Database.Save(&ride)

	return ride
}
