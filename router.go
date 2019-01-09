package main

import (
	"net/http"
	
	"github.com/julienschmidt/httprouter"
)

func NewRouter(h *Handlers) *httprouter.Router {
	router := httprouter.New()
	router.GET("/healthz", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		healthz(w)
	})
	router.POST("/new", h.Create)
	router.GET("/", h.Index)
	router.GET("/:id", h.Get)
	return router
}
