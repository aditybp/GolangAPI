package database

import (
	"GolangAPI/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect()  {
	connection, err := gorm.Open(mysql.Open("root:@/golangapi"), &gorm.Config{})

	if err != nil {
		panic("tidak bisa koneksi database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Produk{})
}
