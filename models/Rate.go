package models

import (
	"gorm.io/gorm"
)

type Rates struct {
	gorm.Model

	Id uint
	Rate string `json:"rate"`
	Comment string `json:"comment"`
	Songs_Id  int
	Songs Songs `gorm:"foreignKey:Songs_Id"`
	Users_Id int
	Users Users `gorm:"foreignKey:Users_Id"`
}