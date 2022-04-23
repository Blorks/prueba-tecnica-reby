package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reby/app/dtos"
	"reby/app/services"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RideController struct {
	rideService services.RideService
}

func NewRideController(rideService services.RideService) *RideController {
	return &RideController{rideService: rideService}
}

func (c *RideController) RideStartHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(w, err)
		}
	}()

	w.Header().Set("Content-Type", "application/json")

	rideDto := dtos.RideDtoPost{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&rideDto); err != nil {
		fmt.Fprintln(w, http.StatusUnprocessableEntity)
	} else {
		rideDtoGet := c.rideService.InitRide(rideDto)

		response, _ := json.Marshal(rideDtoGet)
		fmt.Fprintln(w, string(response))
	}
}

func (c *RideController) RideFinishHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(w, err)
		}
	}()

	w.Header().Set("Content-Type", "application/json")

	stringId := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(stringId)

	rideDtoGetCost := c.rideService.FinishRide(id)

	response, _ := json.Marshal(rideDtoGetCost)
	fmt.Fprintln(w, string(response))
}
