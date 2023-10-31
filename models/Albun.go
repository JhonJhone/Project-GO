package models

import (
	"gorm.io/gorm"
)

type Albuns struct {
	gorm.Model

	Id    uint `gorm:"primaryKey"`
	Name string `json:"name"`
	Imagem string `json:"imagem"`

}