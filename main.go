package main

import (
	"TWFjaWVqLVJvc2lhaw-/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/fetcher", func(r chi.Router) {
			r.Get("/", handlers.RequestsList)
			r.Post("/", handlers.CreateRequest)
		})
	})
	_ = http.ListenAndServe(":8080", r)
}