package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

type Handlers struct {
	store   *Store
	baseUrl string
}

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse form data and return url
	u := r.FormValue("url")

	// Check that it is a URL
	parsed, err := url.Parse(u)
	if err != nil || parsed.Host == "" || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		log.Print("handlers: invalid url: " + u)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Create a short URL
	id, err := h.store.Create(u)
	if err != nil {
		log.Print("handlers: " + err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Print shortened URL
	fmt.Fprintln(w, h.baseUrl+"/"+id)
}

func (h *Handlers) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Get the short URL
	url, err := h.store.Get(p.ByName("id"))
	if err != nil {
		// Respond to health checks
		if p.ByName("id") == "healthz" {
			healthz(w)
			return
		}
		if err != sql.ErrNoRows {
			log.Print("handlers: " + err.Error())
		}
		http.Error(w, "Not found, or an error occurred, idk.", http.StatusNotFound)
		return
	}

	// Redirect to URL
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *Handlers) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Dump HTML
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(AssetIndex))
}

func NewHandlers(store *Store, baseUrl string) *Handlers {
	return &Handlers{
		store:   store,
		baseUrl: baseUrl,
	}
}
