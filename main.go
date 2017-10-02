package main

import (
	"os"
	"log"
	"net/http"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	listenPort := os.Getenv("PORT")
	if len(listenPort) < 1 {
		listenPort = "8080"
	}
	baseUrl := os.Getenv("BASEURL")
	if len(baseUrl) < 4 {
		baseUrl = "http://localhost:" + listenPort
	}
	postgresString := os.Getenv("POSTGRES")
	if len(postgresString) < 1 {
		postgresString = "postgresql://root@localhost:26257/short?sslmode=disable"
	}

	// Open database connection
	db, err := sql.Open("postgres", postgresString)
	if err != nil {
		log.Fatal(err)
	}
	
	// Create storage abstraction
	store := NewStore(db)
	// Setup handlers
	handlers := NewHandlers(store, baseUrl)
	
	// Bind handlers to paths
	http.HandleFunc("/new", handlers.Create)
	http.HandleFunc("/", handlers.Get)

	// Listen
	log.Fatal(http.ListenAndServe(":" + listenPort, nil))
}

