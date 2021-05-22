package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/rohaquinlop/Trainin-exercise/backend/loaddata"

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

	r.Route("/load", func(r chi.Router) {
		r.With(paginate).Post("/", loadAllData) //Post /load
		r.With(paginate).Post("/{date}", loadDateData)
	})

	//graphql server
	r.Handle("/graphql", srv)

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

func loadAllData(w http.ResponseWriter, r *http.Request) {
	loaddata.LoadDataDB("")
	_, err := exec.Command("/bin/sh", "../removeFiles.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
}

func loadDateData(w http.ResponseWriter, r *http.Request) {
	loaddata.LoadDataDB(chi.URLParam(r, "date"))

	_, err := exec.Command("/bin/sh", "../removeFiles.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
}

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}
