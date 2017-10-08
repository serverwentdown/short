package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var num int
var port int
var baseUrl string
var postgres string

func init() {
	flag.IntVar(&num, "num", 4, "number of characters in shortened url")
	flag.IntVar(&port, "port", 8080, "listen on port")
	flag.StringVar(&baseUrl, "baseurl", "localhost:port", "baseurl URL of short links")
	flag.StringVar(&postgres, "postgres", "postgresql://root@localhost:26257/short?sslmode=disable", "postgres string")
}

func main() {
	// Parse commandline flags
	flag.Parse()
	if baseUrl == "localhost:port" {
		baseUrl = fmt.Sprintf("localhost:%d", port)
	}

	// Open database connection
	db, err := sql.Open("postgres", postgres)
	if err != nil {
		log.Fatal(err)
	}

	// Create storage abstraction
	store := NewStore(db, num)
	// Setup handlers
	handlers := NewHandlers(store, baseUrl)
	// Create router
	router := NewRouter(handlers)

	// Listen
	log.Println("main: Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
