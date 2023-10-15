package models

import (
	"github.com/jinzhu/gorm"
)

type Rates struct {
	gorm.Model

	Id    int `json:"id"`
	Songs_Id  int
	Users_Id int
	Rate string `json:"rate"`
	Comment string `json:"comment"`
}