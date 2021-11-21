package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project/app"
	"project/databse"
	"project/models"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	var products []models.Product

	databse.Connector.Find(&products)
	app.Response(rw, 200)
	json.NewEncoder(rw).Encode(products)
}

func Show(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var product models.Product

	databse.Connector.First(&product, id)
	app.Response(rw, 200)
	json.NewEncoder(rw).Encode(product)
}
