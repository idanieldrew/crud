package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"project/databse"
	"project/models"
	"project/routes"
)

func main() {
	initDb()
	router := mux.NewRouter()
	routes.HandleRoute(router)

	fmt.Println("http://localhost:3000 already exist")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func initDb() {
	c := databse.Conf{
		User:     "daniel",
		Database: "crud",
		Password: "Password123#@!",
	}

	ConnectStr := databse.ConnectToDb(c)
	err := databse.ConnectToMysql(ConnectStr)
	if err != nil {
		panic(err.Error())
	}
	databse.Migrate(&models.Product{})
}
