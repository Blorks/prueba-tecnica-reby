package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reby/cmd/server"
	"runtime"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:toor@tcp(localhost:3333)/prueba-tecnica-reby?charset=utf8mb4&parseTime=True&loc=Local"
const httpPort = 8080

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type ServerConfig struct {
	Server   *server.Server
	ServerUp bool
}

func (s *ServerConfig) SetupServer(includeRide bool) {
	dbConn, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	server := server.Server{
		DBConn: dbConn,
		Port:   httpPort,
	}

	s.ServerUp = true
	s.Server = &server

	s.LoadSQLFiles(includeRide)

	fmt.Println("Servidor levantado en el puerto:", server.Port)

	go server.Start()

	time.Sleep(1 * time.Millisecond) //Wait for the server
}

func (s *ServerConfig) LoadSQLFiles(includeRide bool) {
	LoadSQLFile(s.Server.DBConn, filepath.Join(basepath, "database_drop.sql"))
	LoadSQLFile(s.Server.DBConn, filepath.Join(basepath, "database_create.sql"))
	LoadSQLFile(s.Server.DBConn, filepath.Join(basepath, "data_create.sql"))

	if includeRide {
		LoadSQLFile(s.Server.DBConn, filepath.Join(basepath, "data_add_ride.sql"))
	}
}

func (s *ServerConfig) ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Server.Router.ServeHTTP(rr, req)

	return rr
}
