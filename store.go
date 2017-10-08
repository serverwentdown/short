package main

import (
	"database/sql"
	"log"
)

type Store struct {
	db  *sql.DB
	num int
}

func (s *Store) Create(url string) (id string, err error) {
	id = GenerateID(s.num)
	var existing string
	err = s.db.QueryRow(`
		SELECT url FROM links WHERE id = $1
	`, id).Scan(&existing)
	if err == nil {
		log.Print("store: collision occurred on " + id + ", regenerating...")
		return s.Create(url)
	}
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	err = s.Add(id, url)
	return id, err
}

func (s *Store) Add(id, url string) (err error) {
	_, err = s.db.Exec(`
		INSERT INTO links (id, url) VALUES($1, $2)
	`, id, url)
	return err
}

func (s *Store) Get(id string) (url string, err error) {
	err = s.db.QueryRow(`
		SELECT url FROM links WHERE id = $1
	`, id).Scan(&url)
	return url, err
}

func NewStore(db *sql.DB, num int) *Store {
	// Check for table and initialise if necessary
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS links (
			id VARCHAR(24) PRIMARY KEY,
			url VARCHAR(65535)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	return &Store{db, num}
}
