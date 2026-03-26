package handlers

import (
	"book/models"
	"encoding/json"
	"net/http"
)

var categories = make(map[int]models.Category)
var categoryID = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var result []models.Category
	for _, c := range categories {
		result = append(result, c)
	}
	json.NewEncoder(w).Encode(result)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)

	if category.Name == "" {
		http.Error(w, "Name required", http.StatusBadRequest)
		return
	}

	category.ID = categoryID
	categoryID++
	categories[category.ID] = category

	json.NewEncoder(w).Encode(category)
}
