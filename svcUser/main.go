package main

import (
	"encoding/json"
	"net/http"
	"user/application"

	"github.com/gorilla/mux"
)

const servicePort string = `:8080`
const serviceName string = `User API`

func main() {
	r := mux.NewRouter()
	//ping route
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	//api routs
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/register", Register).Methods("POST")

	app := application.New(serviceName, servicePort, r)
	app.Initialize()
	app.Run()
}
