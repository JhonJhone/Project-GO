package models

type Users struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	IsAdm    int    `json:"isadm" db:"isadm" type:"tinyint"`
}