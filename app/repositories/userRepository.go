package repositories

import (
	"reby/app/models"
)

func GetUser(idUser int64) models.User {
	user := models.User{}

	Database.First(&user, idUser)

	return user
}
