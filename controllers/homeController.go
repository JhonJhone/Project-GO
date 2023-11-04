package controllers

import (
	"log"
	"net/http"
	"text/template"

	"Proj-GO/database"
	"Proj-GO/models"

	"github.com/jinzhu/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
    // Open a connection to the database using GORM
    db, err := database.ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
        return
    }
    defer db.Close() // Ensure the database connection is closed when the function exits.

    // Create a slice to hold the songs retrieved from the database
    var songs []models.Songs

    // Use GORM's Find method to retrieve songs from the database
    if err := db.Find(&songs).Error; err != nil {
        log.Fatal(err)
        return // You should handle the error here and return if an error occurs.
    }

    // Render the "Index" template with the retrieved songs
    tmpl.ExecuteTemplate(w, "Index", songs)
}

func Show(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	// Get the ID from the URL parameter
	nId := r.URL.Query().Get("id")

	// Initialize a pointer to a Songs struct
	s := &models.Songs{}

	// Use GORM's First method to find the song with the given ID
	if err := db.First(s, nId).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// Handle the case where no record with the given ID was found
			http.NotFound(w, r)
		} else {
			// Handle other database-related errors
			log.Fatal(err)
		}
	}

	// Render the template with the s struct
	tmpl.ExecuteTemplate(w, "Show", s)

	// Close the database connection
	db.Close()
}

// Função New apenas exibe o formulário para inserir novos dados
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Função Edit, edita os dados
func Edit(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	// Get the ID from the URL parameter
	nId := r.URL.Query().Get("id")

	// Initialize a pointer to a Songs struct
	s := &models.Songs{}

	// Use GORM's First method to find the song with the given ID
	if err := db.First(s, nId).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// Handle the case where no record with the given ID was found
			http.NotFound(w, r)
			db.Close()
			return
		} else {
			// Handle other database-related errors
			log.Fatal(err)
		}
	}

	// Render the template "Edit" with the s struct
	tmpl.ExecuteTemplate(w, "Edit", s)

	// Close the database connection
	db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	// Open a connection to the database using the dbConn() function.
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	if r.Method == "POST" {
		// Retrieve form values from the HTTP request.
		name := r.FormValue("name")
		description := r.FormValue("description")
		author := r.FormValue("author")
		year := r.FormValue("year")
		duration := r.FormValue("duration")

		// Create a new Songs struct with the form values.
		song := models.Songs{
			Name:        name,
			Description: description,
			Author:      author,
			Year:        year,
			Duration:    duration,
		}

		// Use GORM to create a new record in the "songs" table.
		if err := db.Create(&song).Error; err != nil {
			// Handle the error (e.g., log it and return an error response).
			log.Println("Error creating record:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Log a success message to the console.
		log.Println("Valores inseridos com sucesso!")
	}

	// Close the database connection.
	db.Close()

	// Redirect the user to the home page ("/").
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// Função Update, atualiza valores no banco de dados
func Update(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com o banco de dados usando a função: ConnectDB()
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {
		// Pega os campos do formulário
		name := r.FormValue("name")
		description := r.FormValue("description")
		author := r.FormValue("author")
		year := r.FormValue("year")
		duration := r.FormValue("duration")
		id := r.FormValue("id")

		// Create a Songs struct with the updated values
		updatedSong := models.Songs{
			Name:        name,
			Description: description,
			Author:      author,
			Year:        year,
			Duration:    duration,
		}

		// Use GORM to update the record in the "songs" table with the provided ID
		if err := db.Model(&models.Songs{}).Where("id = ?", id).Updates(updatedSong).Error; err != nil {
			// Handle the error (e.g., log it and return an error response).
			log.Println("Error updating record:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Exibe um log com os valores digitados no formulário
		log.Println("Valores atualizados com sucesso!")
	}

	// Encerra a conexão do ConnectDB()
	db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// Função Delete, deleta valores no banco de dados
func Delete(w http.ResponseWriter, r *http.Request) {
	// Abre conexão com banco de dados usando a função: ConnectDB()
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	nId := r.URL.Query().Get("id")

	// Use GORM to delete the record from the "songs" table based on the provided ID
	if err := db.Where("id = ?", nId).Delete(&models.Songs{}).Error; err != nil {
		// Handle the error (e.g., log it and return an error response).
		log.Println("Error deleting record:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Exibe um log indicando que o registro foi deletado com sucesso
	log.Println("Valor deletado com sucesso")

	// Encerra a conexão do ConnectDB()
	db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))
