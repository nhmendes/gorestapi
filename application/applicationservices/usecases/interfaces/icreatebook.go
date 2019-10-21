package interfaces

import "github.com/nhmendes/gorestapi/application/applicationdto"

// ICreateBook : Executes a write action. This action mutates the state of the system.
type ICreateBook interface {
	Execute(book applicationdto.Book)
}
