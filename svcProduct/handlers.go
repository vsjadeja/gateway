package main

import (
	"encoding/json"
	"net/http"
	"product/database"
	"product/entities"
	"strconv"

	"github.com/gorilla/mux"
)

func getById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	var product entities.Product
	database.Instance.First(&product, id)
	sendWrappedResponse(w, http.StatusOK, product)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)

	if err := product.Validate(); err != nil {
		sendWrappedResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	database.Instance.Create(&product)
	sendWrappedResponse(w, http.StatusCreated, product)
}

func sendWrappedResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
