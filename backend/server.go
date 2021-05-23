package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		r.Post("/", loadAllData) //loadData endpoint
		r.Post("/{date}", loadDateData)
	})

	r.Route("/buyers", func(r chi.Router) {
		r.Get("/", showAllBuyers) //Buyers endpoint
		r.Get("/{id}", showInfoBuyer)
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

func showAllBuyers(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/graphql?query=query+showBuyers{buyers{id%20name%20age}}")

	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var raw map[string]interface{}

	err = json.Unmarshal(body, &raw)

	if err != nil {
		log.Fatal(err)
	}

	out, _ := json.Marshal(raw)

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func showInfoBuyer(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	log.Println(userID)
	//http://localhost:8080/graphql?query=query+findBuyer($id: String!){buyerId(id:$id){id name age}}&variables={"id" : "705d7d9b"}
	var getURL = fmt.Sprintf(`http://localhost:8080/graphql?query={buyerId(id:"%s"){id,name,age},buyerTransactions(id:"%s"){id,ip}}`, string(userID), string(userID))
	resp, err := http.Get(getURL)

	if err != nil {
		log.Fatal("resp error: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var raw map[string]interface{}

	err = json.Unmarshal(body, &raw)

	if err != nil {
		log.Fatal(err)
	}

	out, _ := json.Marshal(raw)

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
