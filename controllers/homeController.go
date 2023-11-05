package controllers

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

	"Proj-GO/database"
	"Proj-GO/models"
)

func Index(c *fiber.Ctx) {
	
    db, err := database.ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
        return
    }
    defer db.Close()

    var songs []models.Songs
    if err := db.Find(&songs).Error; err != nil {
        log.Fatal(err)
        return
    }

	c.Render("Index.html", fiber.Map{
        "Songs": songs,
    });
}

func Show(c *fiber.Ctx) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	nId := c.Query("id")

	s := &models.Songs{}
	if err := db.First(s, nId).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.Status(fiber.StatusNotFound).Send("Not Found")
		} else {
			log.Fatal(err)
			c.Status(fiber.StatusInternalServerError).Send("Internal Server Error")
		}
	}

	c.Render("Show", s)
	db.Close()
}

func New(c *fiber.Ctx) {
	c.Render("New", nil)
}

func Edit(c *fiber.Ctx) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	nId := c.Query("id")

	s := &models.Songs{}
	if err := db.First(s, nId).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.Status(fiber.StatusNotFound).Send("Not Found")
			db.Close()
			return
		} else {
			log.Fatal(err)
			c.Status(fiber.StatusInternalServerError).Send("Internal Server Error")
		}
	}

	c.Render("Edit", s)
	db.Close()
}

func Insert(c *fiber.Ctx) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	if c.Method() == "POST" {
		name := c.FormValue("name")
		description := c.FormValue("description")
		author := c.FormValue("author")
		year := c.FormValue("year")
		duration := c.FormValue("duration")

		song := models.Songs{
			Name:        name,
			Description: description,
			Author:      author,
			Year:        year,
			Duration:    duration,
		}

		if err := db.Create(&song).Error; err != nil {
			log.Println("Error creating record:", err)
			c.Status(fiber.StatusInternalServerError).Send("Internal Server Error")
			return
		}

		log.Println("Valores inseridos com sucesso!")
	}

	db.Close()
	c.Redirect("/", fiber.StatusMovedPermanently)
}

func Update(c *fiber.Ctx) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	if c.Method() == "POST" {
		name := c.FormValue("name")
		description := c.FormValue("description")
		author := c.FormValue("author")
		year := c.FormValue("year")
		duration := c.FormValue("duration")
		id := c.FormValue("id")

		updatedSong := models.Songs{
			Name:        name,
			Description: description,
			Author:      author,
			Year:        year,
			Duration:    duration,
		}

		if err := db.Model(&models.Songs{}).Where("id = ?", id).Updates(updatedSong).Error; err != nil {
			log.Println("Error updating record:", err)
			c.Status(fiber.StatusInternalServerError).Send("Internal Server Error")
			return
		}

		log.Println("Valores atualizados com sucesso!")
	}

	db.Close()
	c.Redirect("/", fiber.StatusMovedPermanently)
}

func Delete(c *fiber.Ctx) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	nId := c.Query("id")

	if err := db.Where("id = ?", nId).Delete(&models.Songs{}).Error; err != nil {
		log.Println("Error deleting record:", err)
		c.Status(fiber.StatusInternalServerError).Send("Internal Server Error")
		return
	}

	log.Println("Valor deletado com sucesso")
	db.Close()
	c.Redirect("/", fiber.StatusMovedPermanently)
}
