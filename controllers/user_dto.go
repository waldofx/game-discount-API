package controllers

import (
	"discount-api/models"
	"time"
)

type UserRequest struct {
	ID		 		int			`gorm:"primaryKey" json:"id"`
	Name     		string 		`json:"name"`
	Email      	string 		`json:"email"`
}

func (req *UserRequest) toModel() models.User{
	return models.User{
		ID: req.ID,
		Name: req.Name,
		Email: req.Email,
	}
}

type UserResponse struct{
	ID       		int    		`json:"id"`
	Name      		string 		`json:"name"`
	Email      	string 		`json:"email"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
	DeletedAt 		time.Time 	`json:"deleted_at"`
}

func newResponse2(modelUsers models.User) UserResponse{
	return UserResponse{
		ID: modelUsers.ID,
		Name: modelUsers.Name,
		Email: modelUsers.Email,
		CreatedAt: modelUsers.CreatedAt, 
		UpdatedAt: modelUsers.UpdatedAt,
		DeletedAt: modelUsers.DeletedAt.Time,
	}
}

func newResponseArray2(modelUsers []models.User) []UserResponse{
	var response []UserResponse

	for _, val := range modelUsers{
		response = append(response, newResponse2(val))
	}
	return response
}