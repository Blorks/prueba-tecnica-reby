package repositories

import (
	"reby/app/models"

	"gorm.io/gorm"
)

type VehicleRepository struct {
	conn *gorm.DB
}

func NewVehicleRepository(conn *gorm.DB) *VehicleRepository {
	return &VehicleRepository{conn: conn}
}

/**
Get from the database the vehicle whose primary key matches the id passed as a parameter
*/
func (r *VehicleRepository) GetVehicle(idVehicle int) (models.Vehicle, error) {
	var vehicle models.Vehicle

	if err := r.conn.Find(&vehicle, idVehicle).Error; err != nil {
		return models.Vehicle{}, err
	}

	return vehicle, nil
}
