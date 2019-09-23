package main

import (
	"TWFjaWVqLVJvc2lhaw-/handlers"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/api", func(r chi.Router) {
		r.Get("/fetcher", handlers.RequestsList)
		r.Post("/fetcher", handlers.CreateRequest)
	})
	http.ListenAndServe(":8080", r)
}

func rr(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title")))
}