package main

import (
	"fmt"
	"reby/cmd/server"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "reby:reby@tcp(34.140.134.87:3306)/prueba-tecnica-reby?charset=utf8mb4&parseTime=True&loc=Local"
const httpPort = 8080

func main() {
	dbConn, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	s := server.Server{
		DBConn: dbConn,
		Port:   httpPort,
	}

	fmt.Println("Servidor levantado en el puerto:", s.Port)

	s.Start()
}
