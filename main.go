package main

import (
	"log"

	"Proj-GO/database"
	"Proj-GO/models"

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

	app.Get("/", func(c *fiber.Ctx) error {
		db, err := database.ConnectDB()
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
			return err
		}
		defer db.Close()
	
		var songs []models.Songs
		if err := db.Find(&songs).Error; err != nil {
			log.Fatal(err)
			return err
		}
	
		err = c.Render("Index", songs)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			return err
		}
	
		return nil
	})
	

	log.Println("Server started on: http://localhost:9000")

	app.Listen(":9000")
}

