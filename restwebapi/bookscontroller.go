package restwebapi

import (
	"encoding/json"
	"net/http"

	"github.com/nhmendes/gorestapi/application/applicationdto"
	"github.com/nhmendes/gorestapi/application/applicationservices/usecases/implementation"

	//"github.com/graphql-go/graphql"

	"github.com/gin-gonic/gin"
)

// GetBooks - Get all books
func GetBooks(c *gin.Context) {
	/*
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
	*/
	// ---------------------

	c.Writer.Header().Set("Content-Type", "application/json")

	getBooks := implementation.NewGetBooks()
	result := getBooks.Execute()

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

	getBookById := implementation.NewGetBookByID()
	result, err := getBookById.Execute(c.Params.ByName("id"))

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
	err := json.NewDecoder(c.Request.Body).Decode(&book)

	if err != nil {
		//c.Writer.WriteHeader(http.StatusInternalServerError)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "message": "error!"})
	}

	createBook := implementation.NewCreateBook()
	bookId, err := createBook.Execute(book)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "message": "error!"})
	}

	//c.Writer.WriteHeader(http.StatusCreated)
	c.JSON(
		http.StatusCreated,
		gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": bookId})
}

// UpdateBook - Update book
func UpdateBook(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")

	var book applicationdto.Book
	_ = json.NewDecoder(c.Request.Body).Decode(&book)
	book.ID = c.Params.ByName("id")

	updateBook := implementation.NewUpdateBook()
	updateBook.Execute(book)

	c.Writer.WriteHeader(http.StatusNoContent)
}

// DeleteBook - Delete book
func DeleteBook(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")

	deleteBook := implementation.NewDeleteBook()
	deleteBook.Execute(c.Params.ByName("id"))

	c.Writer.WriteHeader(http.StatusNoContent)
}
