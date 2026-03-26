package handlers

import (
	"book/models"
	"encoding/json"
	"net/http"
)

var authors = make(map[int]models.Author)
var authorID = 1

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var result []models.Author
	for _, a := range authors {
		result = append(result, a)
	}
	json.NewEncoder(w).Encode(result)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)

	if author.Name == "" {
		http.Error(w, "Name required", http.StatusBadRequest)
		return
	}

	author.ID = authorID
	authorID++
	authors[author.ID] = author

	json.NewEncoder(w).Encode(author)
}
