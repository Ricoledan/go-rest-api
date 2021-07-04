package main

import (
	"go-rest-api/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterProfileRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}
