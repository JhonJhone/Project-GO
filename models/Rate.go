package models

import (
	"gorm.io/gorm"
)

type Rates struct {
	gorm.Model

	Id    uint `gorm:"primaryKey"`
	Songs_Id  int
	Users_Id int
	Rate string `json:"rate"`
	Comment string `json:"comment"`

}