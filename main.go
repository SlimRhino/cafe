package main

import (
	"log"
	"net/http"

	"./database"
	"./store"
	"github.com/julienschmidt/httprouter"
)

func init() {
	database.Connect("mongodb://admin:i3540core@ds143242.mlab.com:43242/sweet-cafe", "sweet-cafe")
	database.GetProductCollections()
}

func main() {
	router := httprouter.New()

	router.GET("/store/:collection/:id", store.Get)
	router.POST("/store/:collection", store.Add)
	log.Println("starting server...")
	http.ListenAndServe(":3000", router)
}
