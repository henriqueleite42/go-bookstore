package routes

import (
	"github.com/gorilla/mux"

	"go-bookstore/pkg/controllers"
)

func RegisterBookStoreRoutes (router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")

	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")

	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")

	router.HandleFunc("/books/{bookId}", controllers.EditBook).Methods("PUT")

	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}