package main

import (
	"Proj-GO/models"
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
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

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "melodymeter"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db

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
    // Abre a conexão com o banco de dados utilizando a função dbConn()
    db := dbConn()

    // Create a slice to hold the songs retrieved from the database.
    var songs []models.Songs

    // Query the database to retrieve songs.
    selDB, err := db.Query("SELECT * FROM songs ORDER BY id DESC")
    if err != nil {
        log.Fatal(err)
    }
    defer selDB.Close()

    // Iterate through the query results and populate the 'songs' slice.
    for selDB.Next() {
        var song models.Songs
        err := selDB.Scan(&song.Id, &song.Name, &song.Description, &song.Author, &song.Year, &song.Duration)
        if err != nil {
            log.Fatal(err)
        }
        songs = append(songs, song)
    }

    // Abre a página Index e exibe todos os registrados na tela
    database.ExecuteTemplate(w, "Index", songs)

    // Fecha a conexão
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    // Get the ID from the URL parameter
    nId := r.URL.Query().Get("id")

    // Use the ID to query the database and handle errors
    selDB, err := db.Query("SELECT * FROM songs WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    defer selDB.Close()

    // Initialize a pointer to a Songs struct
    s := &models.Songs{}

    // Check if there is a result (one row) and scan it
    if selDB.Next() {
        var id int
        var name, description, author, year, duration string

        // Scan the result into the s struct
        err = selDB.Scan(&id, &name, &description, &author, &year, &duration)
        if err != nil {
            panic(err.Error())
        }

        // Populate the s struct with the result
        s.Id = id
        s.Name = name
        s.Description = description
        s.Author = author
        s.Year = year
        s.Duration = duration
    } else {
        // Handle the case where no record with the given ID was found
        http.NotFound(w, r)
        return
    }

    // Render the template with the s struct
    database.ExecuteTemplate(w, "Show", s)

    // Close the database connection
    defer db.Close()
}


// Função New apenas exibe o formulário para inserir novos dados
func New(w http.ResponseWriter, r *http.Request) {
	database.ExecuteTemplate(w, "New", nil)
}

// Função Edit, edita os dados
func Edit(w http.ResponseWriter, r *http.Request) {
	// Open a database connection
	db := dbConn()

	// Get the ID from the URL parameter
	nId := r.URL.Query().Get("id")

	// Query the database to retrieve the song with the given ID
	selDB, err := db.Query("SELECT * FROM songs WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()

	// Initialize a pointer to a Songs struct
	s := &models.Songs{}

	// Check if there is a result (one row) and scan it
	if selDB.Next() {
		var id int
		var name, description, author, year, duration string

		// Scan the result into the s struct
		err = selDB.Scan(&id, &name, &description, &author, &year, &duration)
		if err != nil {
			panic(err.Error())
		}

		// Populate the s struct with the result
		s.Id = id
		s.Name = name
		s.Description = description
		s.Author = author
		s.Year = year
		s.Duration = duration
	} else {
		// Handle the case where no record with the given ID was found
		http.NotFound(w, r)
		return
	}

	// Render the template "Edit" with the s struct
	database.ExecuteTemplate(w, "Edit", s)

	// Close the database connection
	defer db.Close()
}


func Insert(w http.ResponseWriter, r *http.Request) {
	// Open a connection to the database using the dbConn() function.
	db := dbConn()

	// Check if the HTTP request method is "POST."
	if r.Method == "POST" {
		// Retrieve form values from the HTTP request.
		name := r.FormValue("name")
		description := r.FormValue("description")
		author := r.FormValue("author")
		year := r.FormValue("year")
		duration := r.FormValue("duration")

		// Prepare an SQL statement for inserting data into the "songs" table.
		insForm, err := db.Prepare("INSERT INTO songs(name, description, author, year, duration) VALUES(?,?,?,?,?)")
		if err != nil {
			// Handle the error (e.g., log it and return an error response).
			log.Println("Error preparing SQL statement:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Execute the prepared statement to insert the data into the table.
		_, err = insForm.Exec(name, description, author, year, duration)
		if err != nil {
			// Handle the error (e.g., log it and return an error response).
			log.Println("Error executing SQL statement:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Log a success message to the console.
		log.Println("Valores inseridos com sucesso!")
	}

	// Close the database connection.
	defer db.Close()

	// Redirect the user to the home page ("/").
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}


// Função Update, atualiza valores no banco de dados
func Update(w http.ResponseWriter, r *http.Request) {

	// Abre a conexão com o banco de dados usando a função: dbConn()
	db := dbConn()

	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		name := r.FormValue("name")
		description := r.FormValue("description")
		author := r.FormValue("author")
		year := r.FormValue("year")
		duration := r.FormValue("duration")
		id := r.FormValue("id")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("UPDATE songs SET name=?, description=?, author=?, year=?, duration=? WHERE songs.id=?")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulário com a SQL tratada e verifica erros
		insForm.Exec(name, description, author, year, duration, id)

		// Exibe um log com os valores digitados no formulario
		log.Println("Valores atualizados com sucesso!")
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// Função Delete, deleta valores no banco de dados
func Delete(w http.ResponseWriter, r *http.Request) {

	// Abre conexão com banco de dados usando a função: dbConn()
	db := dbConn()

	nId := r.URL.Query().Get("id")

	// Prepara a SQL e verifica errors
	delForm, err := db.Prepare("DELETE FROM songs WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do form com a SQL tratada e verifica errors
	delForm.Exec(nId)

	// Exibe um log com os valores digitados no form
	log.Println("Valor deletado com sucesso")

	// Encerra a conexão do dbConn()
	defer db.Close()

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
