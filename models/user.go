package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID		 		int			`gorm:"primaryKey" json:"id"`
	Name     		string 		`json:"name"`
	Email			string		`json:"email"`
	gorm.Model 					//`json:"-"`
}