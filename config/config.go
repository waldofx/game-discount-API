package config

import (
	"discount-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	connectionString := "root:@tcp(0.0.0.0:3306)/project_temp?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	//DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Game{})
}