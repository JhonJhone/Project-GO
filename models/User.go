package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	
	Id uint
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password []byte  `json:"password"`
	IsAdm int `json:"isadm" db:"isadm" type:"tinyint"`
}