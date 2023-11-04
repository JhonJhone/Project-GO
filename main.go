package main

import (
	"log"

	"Proj-GO/database"

	"Proj-GO/routes"

	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    database.ConnectDB()

    app := fiber.New()
    routes.Setup(app)

	log.Println("Server started on: http://localhost:9000")

	app.Listen(":9000", nil)
}

