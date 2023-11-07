package main

import (
	"log"

	"Proj-GO/database"
	"Proj-GO/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    database.ConnectDB()

	engine := html.New("./templates", ".html")
	

    app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Setup(app)

	log.Println("Server started on: http://localhost:9000")

	app.Listen(":9000")
}

