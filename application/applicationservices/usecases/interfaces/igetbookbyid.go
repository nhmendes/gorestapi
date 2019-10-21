package interfaces

import "github.com/nhmendes/gorestapi/application/applicationdto"

// IGetBookByID : Executes a read-only action. This action MUST NOT mutate the state of the system (read-only).
type IGetBookByID interface {
	Execute(id string) applicationdto.Book
}
