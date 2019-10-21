package restwebapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nhmendes/gorestapi/application/applicationdto"
	"github.com/nhmendes/gorestapi/application/applicationservices/usecases/implementation"

	"github.com/graphql-go/graphql"

	"github.com/gin-gonic/gin"
)

// GetBooks - Get all books
func GetBooks(c *gin.Context) {

	// Schema
	fields := graphql.Fields{
		"GetBooks": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			book
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}

	// ---------------------

	c.Writer.Header().Set("Content-Type", "application/json")

	getbooks := implementation.NewGetBooks()
	result := getbooks.Execute()

	c.Writer.WriteHeader(http.StatusOK)
	err2 := json.NewEncoder(c.Writer).Encode(result)

	if err2 != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

// GetBook godoc
// @Summary executes the get book by id use case
// @Description gets a book by ID
// @ID string
// @Tags books
// @Accept  json
// @Produce  json
// @Param  id path string true "Book ID"
// @Success 200 {object} Book
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /bottles/{id} [get]
func GetBook(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")

	getbookbyid := implementation.NewGetBookByID()
	result, err := getbookbyid.Execute(c.Params.ByName("id"))

	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
	}

	c.Writer.WriteHeader(http.StatusOK)
	encodeError := json.NewEncoder(c.Writer).Encode(result)

	if encodeError != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

// CreateBook - Add new book
func CreateBook(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")
	var book applicationdto.Book
	_ = json.NewDecoder(c.Request.Body).Decode(&book)

	createbook := implementation.NewCreateBook()
	createbook.Execute(book)

	c.Writer.WriteHeader(http.StatusCreated)
}

// UpdateBook - Update book
func UpdateBook(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")

	var book applicationdto.Book
	_ = json.NewDecoder(c.Request.Body).Decode(&book)
	book.ID = c.Params.ByName("id")

	updatebook := implementation.NewUpdateBook()
	updatebook.Execute(book)

	c.Writer.WriteHeader(http.StatusNoContent)
}

// DeleteBook - Delete book
func DeleteBook(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")

	deletebook := implementation.NewDeleteBook()
	deletebook.Execute(c.Params.ByName("id"))

	c.Writer.WriteHeader(http.StatusNoContent)
}
