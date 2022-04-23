package server

import (
	"fmt"
	"log"
	"net/http"
	"reby/app/controllers"
	"reby/app/repositories"
	"reby/app/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type Server struct {
	Port   int
	DBConn *gorm.DB
	Router *chi.Mux
}

func (s *Server) Start() {
	userRepository := repositories.NewUserRepository(s.DBConn)
	vehicleRepository := repositories.NewVehicleRepository(s.DBConn)
	rideRepository := repositories.NewRideRepository(s.DBConn)

	rideService := services.NewRideService(*userRepository, *vehicleRepository, *rideRepository)

	rideController := controllers.NewRideController(*rideService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/rides", rideController.RideStartHandler)
	r.Post("/rides/{id}/finish", rideController.RideFinishHandler)

	s.Router = r

	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), r); err != http.ErrServerClosed && err != nil {
		log.Fatalf("Error starting http server <%s>", err)
	}
}
