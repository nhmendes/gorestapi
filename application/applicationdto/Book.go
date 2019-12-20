package applicationdto

import "fmt"

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// ShowDetails - prints the book details
func (b Book) ShowDetails() {
	fmt.Println("Book Title: ", b.Title)
}
