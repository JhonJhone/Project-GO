package models

import (
	"github.com/jinzhu/gorm"
)

type Songs struct {
	gorm.Model

	Id    int `json:"id"`
	Name  string `json:"name"`
	Description string `json:"desciption"`
	Author string `json:"author"`
	Year string `json:"year"`
	Duration string `json:"duration"`
}
