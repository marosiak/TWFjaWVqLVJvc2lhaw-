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
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Route("/fetcher", func(r chi.Router) {
			r.Get("/", handlers.RequestsList)
			r.Post("/", handlers.CreateRequest)
			r.Route("/{requestId}", func(r chi.Router) {
				r.Get("/", handlers.RequestDetail)
				r.Delete("/", handlers.DeleteRequest)
				r.Get("/history", handlers.RequestHistory)
			})
		})
	})
	_ = http.ListenAndServe(":8080", r)
}