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

/**
Get from the database the ride whose primary key matches the id passed as a parameter
*/
func (r *RideRepository) GetRide(idRide int) (models.Ride, error) {
	var ride models.Ride

	if err := r.conn.Find(&ride, idRide).Error; err != nil {
		return models.Ride{}, err
	}

	return ride, nil
}

/**
Insert or update a ride. In case of inserting, it does not insert the user and vehicle relations
*/
func (r *RideRepository) SaveRide(ride models.Ride) models.Ride {
	r.conn.Omit(clause.Associations).Save(&ride)

	return ride
}
