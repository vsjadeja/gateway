package main

import (
	"encoding/json"
	"net/http"
	"product/application"

	"github.com/gorilla/mux"
)

const servicePort string = `:8080`
const serviceName string = `Product API`

func main() {
	r := mux.NewRouter()
	//ping route
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	//api routs
	r.HandleFunc("/{id}", getById).Methods("GET")
	r.HandleFunc("/add", createProduct).Methods("POST")

	app := application.New(serviceName, servicePort, r)
	app.Initialize()
	app.Run()
}
