package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	
	Id    uint    `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email *string `json:"email"`
	Password string  `json:"password"`
	IsAdm int `json:"isadm" db:"isadm" type:"tinyint"`
}