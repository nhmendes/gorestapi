package restwebapi

import (
	"encoding/json"
	"net/http"
	"restapi/application/applicationdto"
	"restapi/application/applicationservices/usecases/implementation"

	"github.com/gorilla/mux"
)

// GetBooks - Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var getbooks = implementation.NewGetBooks()
	var result = getbooks.Execute()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// GetBook - Get single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	var getbookbyid = implementation.NewGetBookByID()
	var result = getbookbyid.Execute(params["id"])

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// CreateBook - Add new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book applicationdto.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	var createbook = implementation.NewCreateBook()
	createbook.Execute(book)

	w.WriteHeader(http.StatusCreated)
}

// UpdateBook - Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var book applicationdto.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"]

	var updatebook = implementation.NewUpdateBook()
	updatebook.Execute(book)

	w.WriteHeader(http.StatusNoContent)
}

// DeleteBook - Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var deletebook = implementation.NewDeleteBook()
	deletebook.Execute(params["id"])

	w.WriteHeader(http.StatusNoContent)
}
