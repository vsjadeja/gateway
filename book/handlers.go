package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type book struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func getById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	book := book{Id: id}
	jsonResponse, jsonError := json.Marshal(book)

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
