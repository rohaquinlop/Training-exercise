package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rohaquinlop/Trainin-exercise/backend/graph"
	"github.com/rohaquinlop/Trainin-exercise/backend/graph/generated"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	//Frontend files
	rootDir := http.Dir("../frontend/")
	http.Handle("/", http.FileServer(rootDir))

	//graphql server
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/home", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
