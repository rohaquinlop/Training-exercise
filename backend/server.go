package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rohaquinlop/Trainin-exercise/backend/graph"
	"github.com/rohaquinlop/Trainin-exercise/backend/graph/generated"
)

const defaultPort = "8080"

func getActualDateTime(value string) string {
	if value == "" {
		date := fmt.Sprintf("%d", time.Now().Unix())
		return date
	} else {
		tmp, _ := time.Parse(time.RFC3339, value)
		date := fmt.Sprintf("%d", tmp.Unix())
		return date
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	localDate := getActualDateTime("")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	//Frontend files
	rootDir := http.Dir("../frontend/")
	http.Handle("/", http.FileServer(rootDir))

	//graphql server
	http.Handle("/graphql", srv)

	log.Printf("local unix timestamp: %s\n", localDate)

	log.Printf("connect to http://localhost:%s/home", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
