package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	initDB()
	r := mux.NewRouter()
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
