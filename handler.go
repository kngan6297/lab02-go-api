package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT id, name, price FROM products")
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Name, &p.Price)
		products = append(products, p)
	}
	json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	json.NewDecoder(r.Body).Decode(&p)
	db.QueryRow("INSERT INTO products(name, price) VALUES($1, $2)", p.Name, p.Price)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product created"})
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p Product
	json.NewDecoder(r.Body).Decode(&p)
	db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3", p.Name, p.Price, id)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated"})
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	db.Exec("DELETE FROM products WHERE id=$1", id)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted"})
}
