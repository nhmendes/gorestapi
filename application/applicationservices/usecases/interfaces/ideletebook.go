package interfaces

// IDeleteBook : Executes a write action. This action mutates the state of the system.
type IDeleteBook interface {
	Execute(id string)
}
