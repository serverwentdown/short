package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

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
	// Create router
	router := NewRouter(handlers)

	// Listen
	log.Println("Listening on port " + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, router))
}
