package models

import (
	"gorm.io/gorm"
)

type Game struct {
	ID		 		int			`gorm:"primaryKey" json:"id"`
	Name     		string 		`json:"name"`
	Description     string  	`json:"description"`
	Category      	string 		`json:"category"`
	//SellerID		int    		`json:"client_id"`
	gorm.Model 					//`json:"-"`
}