package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type Handlers struct {
	store *Store
	baseUrl string
}

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	// Parse form, return bad request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	
	u := r.Form.Get("url")

	// Check that it is a URL
	_, err = url.ParseRequestURI(u)
	if err != nil {
		fmt.Fprintln(w, "How to use: /new?url=http://example.com/")
		return
	}

	// Create a short URL
	id, err := h.store.Create(u)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, h.baseUrl + "/" + id)
}

func (h *Handlers) Get(w http.ResponseWriter, r *http.Request) {
	// Get the short URL
	url, err := h.store.Get(r.URL.Path[1:])
	if err != nil {
		http.Error(w, "Not found, or an error occurred, idk.", http.StatusNotFound)
		return
	}

	// Redirect to URL
	http.Redirect(w, r, url, http.StatusFound)
}

func NewHandlers(store *Store, baseUrl string) *Handlers {
	return &Handlers{
		store: store,
		baseUrl: baseUrl,
	}
}
