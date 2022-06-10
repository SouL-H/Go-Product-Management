package main

import (
	. "Product/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starting...")

	r := mux.NewRouter().StrictSlash(true)//Sonuna / eklediğimizde problemi ortadan kaldırır.
	r.HandleFunc("/api/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products", PostProductHandler).Methods("POST")
	r.HandleFunc("/api/products/{id}", PutProductHandler).Methods("PUT")
	r.HandleFunc("/api/products/{id}", DeleteProductHandler).Methods("DELETE")

	http.ListenAndServe(":8080", r)

	defer log.Println("Server ending...")
}
