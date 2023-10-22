package main

import (
	"log"
	"net/http"
	"text/template"

	"Proj-GO/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// type SpotifySearchResponse struct {
//     Tracks struct {
//         Items []struct {
//             Name     string `json:"name"`
//             Artist   string `json:"artist"`
//             Album    string `json:"album"`
//             // Add more fields as needed
//         } `json:"items"`
//     } `json:"tracks"`
// }

func ConnectDB() (*gorm.DB, error) {
    // Load environment variables from the .env file
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    // Get the environment variables
    // user := os.Getenv("root")
    // password := os.Getenv("")
    // hostname := os.Getenv("localhost")
    // port := os.Getenv("3306")
    // dbname := os.Getenv("melodymeter")

    // Define the DB connection string for MySQL
    dsn := "root:@tcp(127.0.0.1:3306)/melodymeter?charset=utf8mb4&parseTime=True&loc=Local"    // Initialize the MySQL database connection
    db, err := gorm.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Do not defer db.Close() here; it should be managed outside this function.

    // Migrate the schema (automatically create tables if they don't exist)
    db.AutoMigrate(&models.Users{})

    return db, nil
}



// func ConnAPI(r *http.Request, w http.ResponseWriter) {
//     accessToken := "11dFghVXANMlKmJXsNCbNl"
// 	if r.Method == "POST" {
//         // Retrieve form values from the HTTP request.
//         search := r.FormValue("search")

//         // Construct the Spotify API request.
//         apiUrl := "https://api.spotify.com/v1/search"
//         searchType := "track"

//         // Create a request with the access token and query parameters.
//         req, err := http.NewRequest("GET", apiUrl, nil)
//         if err != nil {
//             fmt.Println("Error creating request:", err)
//             return
//         }

//         req.Header.Add("Authorization", "Bearer "+accessToken)
//         q := req.URL.Query()
//         q.Add("q", search)
//         q.Add("type", searchType)
//         req.URL.RawQuery = q.Encode()

//         // Make the GET request.
//         client := &http.Client{}
//         resp, err := client.Do(req)
//         if err != nil {
//             fmt.Println("Error making request:", err)
//             return
//         }
//         defer resp.Body.Close()

//         // Read and handle the response.
//         if resp.StatusCode == http.StatusOK {
//             var spotifyResponse SpotifySearchResponse
//             decoder := json.NewDecoder(resp.Body)
//             err := decoder.Decode(&spotifyResponse)
//             if err != nil {
//                 fmt.Println("Error decoding JSON response:", err)
//                 return
//             }

//             // Process the Spotify search results (e.g., display them in your template).
//             // Access them as spotifyResponse.Tracks.Items
//         } else {
//             fmt.Println("Error:", resp.Status)
//         }
//         // Redirect or render the search results in your template.
//     }
// }


func Index(w http.ResponseWriter, r *http.Request) {
    // Open a connection to the database using GORM
    db, err := ConnectDB()
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
    database.ExecuteTemplate(w, "Index", songs)
}



func Show(w http.ResponseWriter, r *http.Request) {
    db, err := ConnectDB()
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
    database.ExecuteTemplate(w, "Show", s)

    // Close the database connection
    db.Close()
}



// Função New apenas exibe o formulário para inserir novos dados
func New(w http.ResponseWriter, r *http.Request) {
	database.ExecuteTemplate(w, "New", nil)
}

// Função Edit, edita os dados
func Edit(w http.ResponseWriter, r *http.Request) {
    db, err := ConnectDB()
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
    database.ExecuteTemplate(w, "Edit", s)

    // Close the database connection
    db.Close()
}



func Insert(w http.ResponseWriter, r *http.Request) {
    // Open a connection to the database using the dbConn() function.
    db, err := ConnectDB()
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
    db, err := ConnectDB()
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
    db, err := ConnectDB()
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


func main() {

	log.Println("Server started on: http://localhost:9000")

	http.HandleFunc("/", Index)
	// http.HandleFunc("/search", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)


	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	http.ListenAndServe(":9000", nil)
}

var database = template.Must(template.ParseGlob("templates/*.html"))
