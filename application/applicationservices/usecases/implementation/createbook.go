package implementation

import (
	"fmt"
	"restapi/application/applicationdto"
	"restapi/application/applicationservices/usecases/interfaces"
)

// CreateBook : new book use case implementation
type CreateBook struct {
	CreateBook interfaces.ICreateBook
}

// NewCreateBook : ctor
func NewCreateBook() *CreateBook {
	return &CreateBook{}
}

// Execute : executes the get all books use case
func (r *CreateBook) Execute(book applicationdto.Book) {
	fmt.Printf("created book %s", book.ID)
}
