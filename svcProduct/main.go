package main

import (
	"encoding/json"
	"log"
	"net/http"
	"product/database"
	"time"

	"github.com/gorilla/mux"
)

const servicePort string = `:8080`

func main() {
	r := mux.NewRouter()
	//ping route
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	//api routs
	r.HandleFunc("/{id}", getById).Methods("GET")
	r.HandleFunc("/add", createProduct).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    servicePort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func init() {
	//user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:root@tcp(productDB:3306)/productDB?charset=utf8mb4&parseTime=True&loc=Local"
	database.Connect(dsn)
	database.Migrate()
}
