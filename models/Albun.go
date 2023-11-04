package models

import (
	"gorm.io/gorm"
)

type Albuns struct {
	gorm.Model

	Name string `json:"name"`
	Imagem string `json:"imagem"`

}