package models

import (
	"gorm.io/gorm"
)

type Albuns struct {
	gorm.Model

	Id uint
	Name string `json:"name"`
	Imagem string `json:"imagem"`

}