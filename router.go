package main

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter(h *Handlers) *httprouter.Router {
	router := httprouter.New()
	router.GET("/healthz", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		healthz(w)
	})
	router.POST("/new", h.Create)
	router.GET("/", h.Index)
	router.GET("/:id", h.Get)
	return router
}
