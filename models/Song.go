package models

import (
	"gorm.io/gorm"
)

type Songs struct {
	gorm.Model

	Name  string `json:"name"`
	Description string `json:"description"`
	Author string `json:"author"`
	Year string `json:"year"`
	Duration string `json:"duration"`
	Albuns_id int
	Albuns Albuns `gorm:"foreignKey:Albuns_id"`
}
