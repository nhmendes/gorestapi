package interfaces

import "restapi/application/applicationdto"

// IGetBooks : Executes a read-only action. This action MUST NOT mutate the state of the system (read-only).
type IGetBooks interface {
	Execute() []applicationdto.Book
}
