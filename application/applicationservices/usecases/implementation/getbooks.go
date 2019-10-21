package implementation

import (
	"github.com/nhmendes/gorestapi/application/applicationdto"
	"github.com/nhmendes/gorestapi/application/applicationservices/usecases/interfaces"
)

// GetBooks : get all books use case implementation
type GetBooks struct {
	GetBooks interfaces.IGetBooks
}

// NewGetBooks : ctor
func NewGetBooks() *GetBooks {
	return &GetBooks{}
}

// Execute : executes the get all books use case
func (r *GetBooks) Execute() []applicationdto.Book {
	var result []applicationdto.Book
	return append(result,
		applicationdto.Book{
			ID:    "1",
			Isbn:  "438227",
			Title: "Book One",
			Author: &applicationdto.Author{
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		applicationdto.Book{
			ID:    "2",
			Isbn:  "8558646",
			Title: "Book Two",
			Author: &applicationdto.Author{
				FirstName: "Jane",
				LastName:  "Fonda",
			},
		},
	)
}
