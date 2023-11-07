package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    connection, err := gorm.Open(mysql.Open("root:@/melodymeter"), &gorm.Config{})
   
    if err != nil{
        panic("Não foi possível realizar a conexão com o banco de dados")
    }

    DB = connection
}