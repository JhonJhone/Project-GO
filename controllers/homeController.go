package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"Proj-GO/database"
	"Proj-GO/models"
)

func Index(c *fiber.Ctx) error {
    db := database.DB

    var songs []models.Songs
    if err := db.Find(&songs).Error; err != nil {
        log.Fatal(err)
        return err
    }

    err := c.Render("Index", songs)
    if err != nil {
        log.Printf("Error rendering template: %v", err)
        return err
    }

    return nil
}



func Show(c *fiber.Ctx) error {
    db := database.DB

    nId := c.Query("id")

    s := &models.Songs{}
    if err := db.First(s, nId).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            c.Status(fiber.StatusNotFound).Send([]byte("Not Found"))
            return nil
        } else {
            log.Fatal(err)
            c.Status(fiber.StatusInternalServerError).Send([]byte("Internal Server Error"))
            return err
        }
    }

    err := c.Render("Show", s)
    if err != nil {
        log.Printf("Error rendering template: %v", err)
        return err
    }

    return nil
}

func New(c *fiber.Ctx) error {
    err := c.Render("New", nil)
    if err != nil {
        log.Printf("Error rendering template: %v", err)
        return err
    }
    return nil
}

func Edit(c *fiber.Ctx) error {
    db := database.DB

    nId := c.Query("id")

    s := &models.Songs{}
    if err := db.First(s, nId).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            c.Status(fiber.StatusNotFound).Send([]byte("Not Found"))
            return nil
        } else {
            log.Fatal(err)
            c.Status(fiber.StatusInternalServerError).Send([]byte("Internal Server Error"))
            return err
        }
    }

    err := c.Render("Edit", s)
    if err != nil {
        log.Printf("Error rendering template: %v", err)
        return err
    }

    return nil
}

func Insert(c *fiber.Ctx) error {
    db := database.DB

    if c.Method() == "POST" {
        name := c.FormValue("name")
        description := c.FormValue("description")
        author := c.FormValue("author")
        year := c.FormValue("year")
        duration := c.FormValue("duration")
        albunsIDStr := c.FormValue("album")

        albunsID, err := strconv.Atoi(albunsIDStr)
        if err != nil {
            log.Println("Error parsing albuns_id:", err)
            c.Status(fiber.StatusBadRequest).Send([]byte("Invalid Input"))
            return err
        }

        song := models.Songs{
            Name:        name,
            Description: description,
            Author:      author,
            Year:        year,
            Duration:    duration,
            AlbunsID:    uint(albunsID),
        }

        if err := db.Create(&song).Error; err != nil {
            log.Println("Error creating record:", err)
            c.Status(fiber.StatusInternalServerError).Send([]byte("Internal Server Error"))
            return err
        }

        log.Println("Valores inseridos com sucesso!")
    }

    c.Redirect("/", fiber.StatusMovedPermanently)

    return nil
}

func Update(c *fiber.Ctx) error {
    db := database.DB

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
            c.Status(fiber.StatusInternalServerError).Send([]byte("Internal Server Error"))
            return err
        }

        log.Println("Valores atualizados com sucesso!")
    }

    c.Redirect("/", fiber.StatusMovedPermanently)

    return nil
}

func Delete(c *fiber.Ctx) error {
    db := database.DB

    nId := c.Query("id")

    if err := db.Where("id = ?", nId).Delete(&models.Songs{}).Error; err != nil {
        log.Println("Error deleting record:", err)
        c.Status(fiber.StatusInternalServerError).Send([]byte("Internal Server Error"))
        return err
    }

    log.Println("Valor deletado com sucesso")

    c.Redirect("/", fiber.StatusMovedPermanently)

    return nil
}
