package implementation

import (
	"fmt"

	"github.com/nhmendes/gorestapi/application/applicationservices/usecases/interfaces"
)

// DeleteBook : new book use case implementation
type DeleteBook struct {
	DeleteBook interfaces.IDeleteBook
}

// NewDeleteBook : ctor
func NewDeleteBook() *DeleteBook {
	return &DeleteBook{}
}

// Execute : executes the delete books use case
func (r *DeleteBook) Execute(id string) {
	fmt.Printf("deleted book %s", id)
}
