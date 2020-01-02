package tests

import (
	"github.com/nhmendes/gorestapi/application/applicationdto"
	"github.com/nhmendes/gorestapi/application/applicationservices/usecases/implementation"
	"testing"
)

func TestExecute(t *testing.T) {

	b := new(applicationdto.Book)
	b.ID = "1"
	b.ShowDetails()

	var book applicationdto.Book
	book = applicationdto.Book{ID: "id", Isbn: "isbn", Title: "title", Author: nil}

	expected := book.ID

	createBook := implementation.NewCreateBook()
	actual, err := createBook.Execute(book)

	//getbookbyid := implementation.NewGetBookByID()
	//result, err := getbookbyid.Execute(book.ID)

	if err != nil {
		t.Error("error")
	}

	if expected != actual {
		t.Error("error")
	}
}
