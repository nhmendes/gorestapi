package interfaces

import "github.com/nhmendes/restapi/application/applicationdto"

// IUpdateBook : Executes a write action. This action mutates the state of the system.
type IUpdateBook interface {
	Execute(book applicationdto.Book)
}
