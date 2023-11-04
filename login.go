package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)



func Register(r *http.Request, w http.ResponseWriter) {
	database.ExecuteTemplate(w, "Register", nil)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {

}

func Login(r *http.Request, w http.ResponseWriter) {
	database.ExecuteTemplate(w, "Login", nil)
}

func Validate(w http.ResponseWriter, r *http.Request) {
	
}
