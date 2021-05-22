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

	FileServer(r)

	//graphql server
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/home", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func FileServer(router *chi.Mux) {
	root := "../frontend/"
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}
