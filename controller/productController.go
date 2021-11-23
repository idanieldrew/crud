package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"project/app"
	"project/databse"
	"project/models"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	var products []models.Product

	databse.Connector.Find(&products)
	app.Response(rw, 200, "")
	json.NewEncoder(rw).Encode(products)
}

func Show(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var product models.Product

	databse.Connector.First(&product, id)
	app.Response(rw, 200, "")
	json.NewEncoder(rw).Encode(product)
}

func Store(rw http.ResponseWriter, r *http.Request) {
	// get request
	body, _ := ioutil.ReadAll(r.Body)

	var product models.Product

	// store body in product
	err := json.Unmarshal(body, &product)
	if err != nil {
		fmt.Println("can't unmarshal body:", err, "")
	}

	// store body in table(product)
	databse.Connector.Create(&product)
	app.Response(rw, 201, "successfully create.")
	json.NewEncoder(rw).Encode(product)
}

func Update(rw http.ResponseWriter, r *http.Request) {
	// get id's
	v := mux.Vars(r)
	id := v["id"]

	body, _ := ioutil.ReadAll(r.Body)

	var product models.Product

	err := json.Unmarshal(body, &product)

	if err != nil {
		fmt.Println("can't unmarshal body:", err)
	}
	databse.Connector.Model(&product).Where("id =? ", id).Update(&product)
	app.Response(rw, 200, "successfully update.")
	json.NewEncoder(rw).Encode(product)
}

func Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product models.Product

	databse.Connector.Model(&product).Delete(&product, "id =?", id)
	app.Response(rw, 204, "successfully delete.")
}
