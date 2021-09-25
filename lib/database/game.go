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

func GetGameByID(id string) (*models.Game, error) {
	var game models.Game

	if err := config.DB.First(&game, id).Error; err != nil {
		return &models.Game{}, err
	}
	return &game, nil
}

func PostGame(game models.Game) (*models.Game, error){
	if err := config.DB.Save(&game).Error; err != nil{
		return &models.Game{}, err
	}
	return &game, nil
}

func UpdateGame(id string, gameData models.Game) (*models.Game, error){
	var game models.Game

	if err := config.DB.First(&game, id).Updates(&gameData).Error; err != nil{
		return &models.Game{}, err
	}
	return &game, nil
}

func DeleteGame(id string, game models.Game) (string, error){
	if err := config.DB.Unscoped().Delete(&game).Error; err != nil{
		return "", err
	}
	return "deleted", nil
}