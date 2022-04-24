package services

import (
	"reby/app/models"
	"reby/app/services"
	"reby/test/unit/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_Init_Ride(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepositoryMock := mocks.NewMockUserRepositoryInterface(ctrl)

	user := models.User{}
	user.Constructor(1, "test", "test@test.com", 100)

	userRepositoryMock.EXPECT().GetUser(1).Return(user).Times(1)

	rideService := services.NewRideService(userRepositoryMock, nil, nil)
}
