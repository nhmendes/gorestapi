// API REST EXAMPLE
//
// This is a example over how to create the api from the source.
//
//     Schemes: http, https
//     Host: localhost:3000
//     Version: 0.1.0
//     basePath: /api
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

package main

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"

	"github.com/nhmendes/gorestapi/restwebapi"

	"github.com/gin-gonic/gin"
	//"github.com/gorilla/mux"
)

var mySigningKey = []byte("mykey")

func isAuthorized(endpoint func(c *gin.Context)) gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {
		if c.Request.Header["Token"] != nil {
			token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				_, _ = fmt.Fprintf(c.Writer, err.Error())
			}

			if token != nil && token.Valid == true {
				endpoint(c)
			}

		} else {
			_, _ = fmt.Fprintf(c.Writer, "not authorized")
		}
	})
}

/*
// APIMiddleware stuff...
func APIMiddleware(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		endpoint(c)
	})
}*/

// Main function
func main() {
	r := gin.Default()

	r.GET("/authorization", restwebapi.GetToken)

	// swagger:operation GET /books Gets all books
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: successful operation
	r.GET("/books", isAuthorized(restwebapi.GetBooks))

	r.GET("/books/{id}", isAuthorized(restwebapi.GetBook))
	r.POST("/books", isAuthorized(restwebapi.CreateBook))
	r.PUT("/books/{id}", isAuthorized(restwebapi.UpdateBook))
	r.DELETE("/books/{id}", isAuthorized(restwebapi.DeleteBook))

	log.Fatal(r.Run(":8001"))
}
