package repositories

import (
	"reby/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{conn: conn}
}

func (r *UserRepository) GetUser(idUser int) (models.User, error) {
	var user models.User

	if err := r.conn.Find(&user, idUser).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
