package repositories

import (
	"reby/app/models"

	"gorm.io/gorm/clause"
)

func GetRide(idRide int) models.Ride {
	ride := models.Ride{}

	Database.Find(&ride, idRide)

	return ride
}

func SaveRide(ride models.Ride) models.Ride {
	Database.Omit(clause.Associations).Save(&ride)

	return ride
}
