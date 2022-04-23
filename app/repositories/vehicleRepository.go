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

func (r *VehicleRepository) GetVehicle(idVehicle int) (models.Vehicle, error) {
	var vehicle models.Vehicle

	if err := r.conn.Find(&vehicle, idVehicle).Error; err != nil {
		return models.Vehicle{}, err
	}

	return vehicle, nil
}
