package main

import (
	"go-rest-api/app"
	"log"
	"net/http"
	"os"
)

func main() {
	app := app.New()

	http.HandleFunc("/", app.Router.ServeHTTP)

	err := http.ListenAndServe(":9000", nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
