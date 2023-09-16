package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"product/database"
	"product/entities"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type errorWrapper struct {
	Errors []string `json:"errors"`
}

func getById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	var product entities.Product
	database.Instance.First(&product, id)
	sendWrappedResponse(w, http.StatusOK, product)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)

	err := validator.New().Struct(product)
	if err != nil {
		sendWrappedResponse(w, http.StatusUnprocessableEntity, wrapValidationError(err))
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

func wrapValidationError(err error) errorWrapper {
	validationErr, ok := err.(validator.ValidationErrors)
	if !ok {
		return errorWrapper{Errors: []string{err.Error()}}
	}

	errorList := make([]string, len(validationErr))
	for _, vErr := range validationErr {
		e := fmt.Sprintf("'%s' has a value of '%v' which does not satisfy '%s'.", vErr.Field(), vErr.Value(), vErr.Tag())
		errorList = append(errorList, e)
	}

	return errorWrapper{Errors: errorList}
}
