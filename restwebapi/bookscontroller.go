package restwebapi

import (
	"encoding/json"
	"net/http"

	"github.com/nhmendes/restapi/application/applicationdto"
	"github.com/nhmendes/restapi/application/applicationservices/usecases/implementation"

	"github.com/gorilla/mux"
)

// GetBooks - Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	getbooks := implementation.NewGetBooks()
	result := getbooks.Execute()

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetBook godoc
// @Summary executes the get book by id use case
// @Description gets a book by ID
// @ID string
// @Tags books
// @Accept  json
// @Produce  json
// @Param  id path string true "Book ID"
// @Success 200 {object} Book
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /bottles/{id} [get]
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	getbookbyid := implementation.NewGetBookByID()
	result, err := getbookbyid.Execute(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	encodeError := json.NewEncoder(w).Encode(result)

	if encodeError != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// CreateBook - Add new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book applicationdto.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	createbook := implementation.NewCreateBook()
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

	updatebook := implementation.NewUpdateBook()
	updatebook.Execute(book)

	w.WriteHeader(http.StatusNoContent)
}

// DeleteBook - Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	deletebook := implementation.NewDeleteBook()
	deletebook.Execute(params["id"])

	w.WriteHeader(http.StatusNoContent)
}
