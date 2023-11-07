package models

import "gorm.io/gorm"

type Songs struct {
    gorm.Model 

	Id uint
    Name        string `json:"name"`
    Description string `json:"description"`
    Author      string `json:"author"`
    Year        string `json:"year"`
    Duration    string `json:"duration"`
    AlbunsID    uint  
    Albuns      Albuns `gorm:"foreignKey:AlbunsID"`
}
