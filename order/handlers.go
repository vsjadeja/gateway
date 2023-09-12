package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type order struct {
	Id   int64  `json:"order_id"`
	Name string `json:"name"`
}

func getById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	order := order{Id: id}
	jsonResponse, jsonError := json.Marshal(order)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	sendWrappedResponse(w, jsonResponse)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	orders := []order{}
	orders = append(orders, order{Id: 1})
	jsonResponse, jsonError := json.Marshal(orders)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	sendWrappedResponse(w, jsonResponse)
}

func sendWrappedResponse(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
