package implementation

import (
	"github.com/nhmendes/restapi/application/applicationdto"
	"github.com/nhmendes/restapi/application/applicationservices/usecases/interfaces"
)

// GetBookByID : get all books use case implementation
type GetBookByID struct {
	GetBookByID interfaces.IGetBookByID
}

// NewGetBookByID : ctor
func NewGetBookByID() *GetBookByID {
	return &GetBookByID{}
}

// Execute : executes the get all books use case
func (r *GetBookByID) Execute(id string) (applicationdto.Book, error) {
	return applicationdto.Book{
		ID:    id, // "1",
		Isbn:  "438227",
		Title: "Book One",
		Author: &applicationdto.Author{
			FirstName: "John",
			LastName:  "Doe",
		},
	}, nil
}
