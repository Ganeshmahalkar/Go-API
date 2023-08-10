package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var db *sql.DB

func main() {
	config := DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "Gtmvijay@1",
		DBName:   "employee",
	}

	connectionString := "host=" + config.Host + " port=" + config.Port + " user=" + config.User +
		" password=" + config.Password + " dbname=" + config.DBName + " sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")

	server := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	log.Println("Server started on localhost:8001")
	//log.Fatal(http.ListenAndServe(":8001", router))
	log.Fatal(server.ListenAndServe())
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	id := 5
	username := "Sudhir"
	email := "sudhir123@gmail.com"

	insertQuery := `
	INSERT INTO users (id, username, email)
	VALUES ($1, $2, $3)
	RETURNING id;
	`

	//err := db.QueryRow(insertQuery, user.ID, user.Username, user.Email).Scan(&user.ID)
	err := db.QueryRow(insertQuery, id, username, email).Scan(&user.ID)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user User

	getUserQuery := `
		SELECT id, username, email
		FROM users
		WHERE id = $1;
	`
	//err := db.QueryRow(getUserQuery, id).Scan(&user.ID, &user.Username, &user.Email)
	err := db.QueryRow(getUserQuery, id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
