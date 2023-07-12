package handlers

import (
	"encoding/json"
	"math/rand"
	"mux/models"
	"net/http"
	"strconv"
)

func CreateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000000))
	//books = append(api.GetMock(), book)
	json.NewEncoder(w).Encode(book)
}
