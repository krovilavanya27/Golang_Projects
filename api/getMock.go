package api

import (
	"mux/models"
)

var books []models.Book

func GetMock() []models.Book {
	books = append(books, models.Book{
		ID:     "1",
		Isbn:   "4424242",
		Title:  "Book x",
		Author: &models.Author{FirstName: "saikiran", LastName: "rallabandi"}})
	return books

}
