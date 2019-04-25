package implementation

import (
	"fmt"

	"github.com/nhmendes/restapi/application/applicationdto"
	"github.com/nhmendes/restapi/application/applicationservices/usecases/interfaces"
)

// UpdateBook : new book use case implementation
type UpdateBook struct {
	UpdateBook interfaces.IUpdateBook
}

// NewUpdateBook : ctor
func NewUpdateBook() *UpdateBook {
	return &UpdateBook{}
}

// Execute : executes the get all books use case
func (r *UpdateBook) Execute(book applicationdto.Book) {
	fmt.Printf("updated book %s", book.ID)
}
