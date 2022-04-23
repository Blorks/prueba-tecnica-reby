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

func Test_Init_Ride_Ok(t *testing.T) {
	if !serverConfig.ServerUp {
		serverConfig.SetupServer(false)
	} else {
		serverConfig.LoadSQLFiles(false)
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

func Test_Finish_Ride_Ok(t *testing.T) {
	if !serverConfig.ServerUp {
		serverConfig.SetupServer(true)
	} else {
		serverConfig.LoadSQLFiles(true)
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
