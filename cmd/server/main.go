package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "reby:reby@tcp(34.140.134.87:3306)/prueba-tecnica-reby?charset=utf8mb4&parseTime=True&loc=Local"
const httpPort = 8080

func main() {
	dbConn, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	s := Server{
		DBConn: dbConn,
		Port:   httpPort,
	}

	s.Start()
}
