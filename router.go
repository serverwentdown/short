package main

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter(h *Handlers) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", h.Index)
	router.POST("/new", h.Create)
	router.GET("/:id", h.Get)
	return router
}
