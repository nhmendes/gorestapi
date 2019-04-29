package main

import (
	"fmt"
	"log"

	"github.com/nhmendes/restapi/restwebapi"

	"github.com/dgrijalva/jwt-go"
	//"github.com/gorilla/mux"

	"github.com/gin-gonic/gin"
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

// Main function
func main() {
	r := gin.Default()
	r.GET("/books", isAuthorized(restwebapi.GetBooks))
	r.GET("/books/{id}", isAuthorized(restwebapi.GetBook))
	r.POST("/books", isAuthorized(restwebapi.CreateBook))
	r.PUT("/books/{id}", isAuthorized(restwebapi.UpdateBook))
	r.DELETE("/books/{id}", isAuthorized(restwebapi.DeleteBook))
	log.Fatal(r.Run(":8001"))
}

// Request sample
// {
// 	"isbn":"4545454",
// 	"title":"Book Three",
// 	"author":{"firstname":"Harry","lastname":"White"}
// }
