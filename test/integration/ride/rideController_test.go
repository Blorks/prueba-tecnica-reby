package ride

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reby/app/dtos"
	"reby/test/integration"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var serverConfig integration.ServerConfig

/* -------------- Init Ride Tests --------------- */
func Test_Init_Ride_Ok(t *testing.T) {
	if !serverConfig.ServerUp {
		serverConfig.SetupServer(false, false, false, false)
	} else {
		serverConfig.LoadSQLFiles(false, false, false, false)
	}

	body := `{"idUser": 1, "idVehicle": 1}`
	request, err := http.NewRequest("POST", "http://localhost:8080/rides", strings.NewReader(body))

	before := time.Now()

	response := serverConfig.ExecuteRequest(request)

	after := time.Now()

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, http.StatusCreated, response.Code)

	responseByte, _ := io.ReadAll(response.Body)

	var dtoResponse dtos.RideDtoGet
	json.Unmarshal(responseByte, &dtoResponse)

	assert.NotEmpty(t, dtoResponse.IdRide)
	assert.True(t, dtoResponse.Created.After(before) && dtoResponse.Created.Before(after))
	assert.Equal(t, dtoResponse.Created, dtoResponse.Finished)
	assert.Equal(t, 1, dtoResponse.IdUser)
	assert.Equal(t, 1, dtoResponse.IdVehicle)
}

func Test_Init_Ride_UserBalanceTooLow(t *testing.T) {
	if !serverConfig.ServerUp {
		serverConfig.SetupServer(false, true, false, false)
	} else {
		serverConfig.LoadSQLFiles(false, true, false, false)
	}

	body := `{"idUser": 1, "idVehicle": 1}`
	request, err := http.NewRequest("POST", "http://localhost:8080/rides", strings.NewReader(body))

	response := serverConfig.ExecuteRequest(request)

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, http.StatusForbidden, response.Code)

	responseByte, _ := io.ReadAll(response.Body)
	responseMsg := string(responseByte)

	assert.Equal(t, "The user's balance is too low to start the ride\n", responseMsg)
}

func Test_Init_Ride_VehicleInUse(t *testing.T) {
	if !serverConfig.ServerUp {
		serverConfig.SetupServer(false, false, true, false)
	} else {
		serverConfig.LoadSQLFiles(false, false, true, false)
	}

	body := `{"idUser": 1, "idVehicle": 1}`
	request, err := http.NewRequest("POST", "http://localhost:8080/rides", strings.NewReader(body))

	response := serverConfig.ExecuteRequest(request)

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, http.StatusForbidden, response.Code)

	responseByte, _ := io.ReadAll(response.Body)
	responseMsg := string(responseByte)

	assert.Equal(t, "The used vehicle is not available at this time\n", responseMsg)
}

/* -------------- Finish Ride Tests --------------- */
func Test_Finish_Ride_Ok(t *testing.T) {
	if !serverConfig.ServerUp {
		serverConfig.SetupServer(true, false, false, false)
	} else {
		serverConfig.LoadSQLFiles(true, false, false, false)
	}

	request, err := http.NewRequest("POST", "http://localhost:8080/rides/1/finish", nil)

	before := time.Now()

	response := serverConfig.ExecuteRequest(request)

	after := time.Now()

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, http.StatusOK, response.Code)

	responseByte, _ := io.ReadAll(response.Body)

	var dtoResponse dtos.RideDtoGetCost
	json.Unmarshal(responseByte, &dtoResponse)

	assert.Equal(t, 1, dtoResponse.IdRide)
	assert.Equal(t, time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local), dtoResponse.Created)
	assert.True(t, dtoResponse.Finished.After(before) && dtoResponse.Finished.Before(after))
	assert.Equal(t, 1, dtoResponse.IdUser)
	assert.Equal(t, 1, dtoResponse.IdVehicle)
	assert.NotEmpty(t, dtoResponse.Cost)
}

func Test_Finish_Ride_RideFinished(t *testing.T) {
	if !serverConfig.ServerUp {
		serverConfig.SetupServer(true, false, false, true)
	} else {
		serverConfig.LoadSQLFiles(true, false, false, true)
	}

	request, err := http.NewRequest("POST", "http://localhost:8080/rides/1/finish", nil)

	response := serverConfig.ExecuteRequest(request)

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, http.StatusForbidden, response.Code)

	responseByte, _ := io.ReadAll(response.Body)
	responseMsg := string(responseByte)

	assert.Equal(t, "The ride that is trying to end is already over\n", responseMsg)
}
