package store

import (
	"encoding/json"
	"log"
	"net/http"

	"../database"
	"../responds"
	"github.com/julienschmidt/httprouter"
)

//Get pruduct
func Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	collection := "prod_" + p.ByName("collection")
	if !database.CollectionExist(collection) {
		responds.JSON(w, http.StatusBadRequest, "invalid collection")
		return
	}
	var store = dao{
		collection: collection,
		db:         database.DB,
	}

	id := p.ByName("id")
	product, err := store.findByID(id)
	if err != nil {
		log.Println(id, "not found in", collection)
		responds.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	responds.JSON(w, http.StatusOK, &product)
	log.Println("successful respond", id, "in", collection, "collection")
}

//Add new product to database
func Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	collection := "prod_" + p.ByName("collection")
	if !database.CollectionExist(collection) {
		responds.JSON(w, http.StatusBadRequest, "invalid collection")
		return
	}

	var store = dao{
		collection: collection,
		db:         database.DB,
	}

	var product Product
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&product)
	if err != nil {
		responds.JSON(w, http.StatusBadRequest, "invalid request data")
		log.Println("error in decode from request body.", err)
		return
	}

	err = store.add(&product)
	if err != nil {
		responds.JSON(w, http.StatusBadRequest, err.Error())
		log.Println("incorrected collection", collection, "in request")
		return
	}

	responds.JSON(w, http.StatusCreated, &product)
	log.Println("create product with id", product.ID, "in", collection, "collection")
}
