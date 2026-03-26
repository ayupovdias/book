package handlers

import (
	"book/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var books = make(map[int]models.Book)
var bookID = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var result []models.Book

	category := r.URL.Query().Get("category")

	for _, book := range books {
		if category != "" {
			if strconv.Itoa(book.CategoryID) != category {
				continue
			}
		}
		result = append(result, book)
	}

	json.NewEncoder(w).Encode(result)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.Title == "" || book.Price <= 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	book.ID = bookID
	bookID++
	books[book.ID] = book

	json.NewEncoder(w).Encode(book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	book, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	var updated models.Book
	json.NewDecoder(r.Body).Decode(&updated)

	updated.ID = id
	books[id] = updated

	json.NewEncoder(w).Encode(updated)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	delete(books, id)

	w.WriteHeader(http.StatusNoContent)
}
