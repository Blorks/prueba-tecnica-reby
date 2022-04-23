package repositories

import (
	"reby/app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RideRepository struct {
	conn *gorm.DB
}

func NewRideRepository(conn *gorm.DB) *RideRepository {
	return &RideRepository{conn: conn}
}

func (r *RideRepository) GetRide(idRide int) (models.Ride, error) {
	var ride models.Ride

	if err := r.conn.Find(&ride, idRide).Error; err != nil {
		return models.Ride{}, err
	}

	return ride, nil
}

func (r *RideRepository) SaveRide(ride models.Ride) models.Ride {
	r.conn.Omit(clause.Associations).Save(&ride)

	return ride
}
