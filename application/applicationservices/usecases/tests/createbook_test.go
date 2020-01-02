package tests

import (
	"testing"
)

func TestExecute(t *testing.T) {

	b := new(applicationdto.Book)
	b.ID = "1"
	b.ShowDetails()

	var book applicationdto.Book
	book = applicationdto.Book{ID: "id", Isbn: "isbn", Title: "title", Author: nil}

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
