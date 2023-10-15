package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string  `json:"password"`
	IsAdm int `json:"isadm" db:"isadm" type:"tinyint"`
}