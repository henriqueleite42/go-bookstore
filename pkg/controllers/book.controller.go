package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(id)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	cb := &models.Book{}
	utils.ParseBody(r, cb)
	b := cb.CreateBook()
	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	bookToUpdate := &models.Book{}

	utils.ParseBody(r, bookToUpdate)

	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(id)

	if (bookToUpdate.Name != "") {
		bookDetails.Name = bookToUpdate.Name
	}

	if (bookToUpdate.Author != "") {
		bookDetails.Author = bookToUpdate.Author
	}

	if (bookToUpdate.Publication != "") {
		bookDetails.Publication = bookToUpdate.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
