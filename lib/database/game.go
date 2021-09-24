package database

import (
	"discount-api/config"
	"discount-api/models"
)

func GetAllGame() (*[]models.Game, error){
	var games []models.Game

	if err := config.DB.Find(&games).Error; err != nil{
		return &[]models.Game{}, err
	}
	return &games, nil
}
