package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	fmt.Printf("Starting server at port 8080\n")

	log.Fatal(http.ListenAndServe(":8080", r))
}