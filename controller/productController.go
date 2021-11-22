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

func Store(rw http.ResponseWriter, r *http.Request) {
	// get request
	body, _ := ioutil.ReadAll(r.Body)

	var product models.Product

	// store body in product
	err := json.Unmarshal(body, &product)
	if err != nil {
		fmt.Println("can't unmarshal body:", err)
	}

	// store body in table(product)
	databse.Connector.Create(&product)
	app.Response(rw, 201)
	json.NewEncoder(rw).Encode(product)
}

func Update(rw http.ResponseWriter, r *http.Request) {
	// get id's
	v := mux.Vars(r)
	id := v["id"]

	body, _ := ioutil.ReadAll(r.Body)

	// get request
	var products []models.Product
	var product models.Product

	databse.Connector.First(&products, id)

	err := json.Unmarshal(body, &product)
	fmt.Println(product)
	if err != nil {
		fmt.Println("can't unmarshal body:", err)
	}
	//databse.Connector.Update(map[string]interface{product.ID, product.Title, product.Body})
}
