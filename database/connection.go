package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() (*gorm.DB, error) {
    // Load environment variables from the .env file
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    dsn := "root:@tcp(127.0.0.1:3306)/melodymeter?charset=utf8mb4&parseTime=True&loc=Local"    // Initialize the MySQL database connection
    db, err := gorm.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    return db, nil
}