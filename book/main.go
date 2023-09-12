package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const servicePort string = `:8080`

func main() {
	r := mux.NewRouter()
	//ping route
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	//api routs
	r.HandleFunc("/{id}", getById).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    servicePort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
