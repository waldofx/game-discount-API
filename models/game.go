package models

import (
	"gorm.io/gorm"
)

type Game struct {
	ID		 		int			`gorm:"primaryKey" json:"id"`
	Name     		string 		`json:"name"`
	Category      	string 		`json:"category"`
	gorm.Model 					//`json:"-"`
}