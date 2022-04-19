package repositories

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "reby:reby@tcp(34.140.134.87:3306)/prueba-tecnica-reby?charset=utf8mb4&parseTime=True&loc=Local"

var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error al establecer la conexión con la base de datos", err)
		panic(err)
	} else {
		fmt.Println("Conexión con la base de datos establecida")
		return db
	}

}()
