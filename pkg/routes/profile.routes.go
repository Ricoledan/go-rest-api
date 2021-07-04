package routes

import (
	"go-rest-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterProfileRoutes = func(router *mux.Router) {
	router.HandleFunc("/profiles/", controllers.GetProfile).Methods("GET")
	router.HandleFunc("/profiles/", controllers.CreateProfile).Methods("POST")
	router.HandleFunc("/profiles/{profileName}", controllers.UpdateProfileByName).Methods("GET")
	router.HandleFunc("/profiles/{profileName}", controllers.DeleteProfileByName).Methods("GET")
}
