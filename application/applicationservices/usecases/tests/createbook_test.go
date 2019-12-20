package tests

import (
	"testing"

	"github.com/nhmendes/gorestapi/application/applicationdto"
	"github.com/nhmendes/gorestapi/application/applicationservices/usecases/implementation"
)

func TestExecute(t *testing.T) {

	var book applicationdto.Book
	book = applicationdto.Book{ID: "id", Isbn: "isbn", Title: "title", Author: nil}

	//book.ShowDetails()

	expected := book.ID

	createbook := implementation.NewCreateBook()
	actual, err := createbook.Execute(book)

	//getbookbyid := implementation.NewGetBookByID()
	//result, err := getbookbyid.Execute(book.ID)

	if err != nil {
		t.Error("erro")
	}

	if expected != actual {
		t.Error("erro")
	}
}
