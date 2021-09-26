package controllers

import (
	"discount-api/models"
	"time"
)

type GameRequest struct {
	ID		 		int			`gorm:"primaryKey" json:"id"`
	Name     		string 		`json:"name"`
	Category      	string 		`json:"category"`
}

func (req *GameRequest) toModel() models.Game{
	return models.Game{
		ID: req.ID,
		Name: req.Name,
		Category: req.Category,
	}
}

type GameResponse struct{
	ID       		int    		`json:"id"`
	Name      		string 		`json:"name"`
	Category      	string 		`json:"category"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
	DeletedAt 		time.Time 	`json:"deleted_at"`
}

func newResponse(modelGames models.Game) GameResponse{
	return GameResponse{
		ID: modelGames.ID,
		Name: modelGames.Name,
		Category: modelGames.Category,
		CreatedAt: modelGames.CreatedAt, 
		UpdatedAt: modelGames.UpdatedAt,
		DeletedAt: modelGames.DeletedAt.Time,
	}
}

func newResponseArray(modelGames []models.Game) []GameResponse{
	var response []GameResponse

	for _, val := range modelGames{
		response = append(response, newResponse(val))
	}
	return response
}