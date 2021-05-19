package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	frontDir := http.Dir("../frontend")

	http.Handle("/", http.FileServer(frontDir))

	log.Println("Server on port 3333")
	log.Fatal(http.ListenAndServe(":3333", nil))
	http.ListenAndServe(":3333", r)
}
