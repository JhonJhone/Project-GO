package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)



func Register(r *http.Request, w http.ResponseWriter) {
	db := dbConn()
}

func Login(r *http.Request, w http.ResponseWriter) {
	
}