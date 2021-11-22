package routes

import (
	"github.com/gorilla/mux"
	"project/controller"
)

func HandleRoute(router *mux.Router) {
	// get all products
	router.HandleFunc("/", controller.Index).Methods("GET")

	// get single product
	router.HandleFunc("/{id}", controller.Show).Methods("GET")

	// add product
	router.HandleFunc("/store", controller.Store).Methods("POST")

	// update product
	router.HandleFunc("/update/{id}", controller.Update).Methods("PUT")
}
