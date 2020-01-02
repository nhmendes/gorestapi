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
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/nhmendes/gorestapi/restwebapi"

	"github.com/gin-gonic/gin"
	//"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

//var mySigningKey = []byte("mykey")
/*
func users(c *gin.Context) {
	/*decoded := c.Request.Header.Get("decoded")
	var user User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)*
}*/

func isAuthorized(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
					}
					return []byte("secret"), nil
				})
				if err != nil {
					_ = json.NewEncoder(c.Writer).Encode(restwebapi.ErrorMsg{Message: err.Error()})
					return
				}
				if token.Valid {
					var user restwebapi.User
					_ = mapstructure.Decode(token.Claims, &user)
					
					name := c.Params.ByName("userId")
					if name != user.Username {
						_ = json.NewEncoder(c.Writer).Encode(restwebapi.ErrorMsg{Message: "Invalid authorization token - Does not match UserID"})
						return
					}

					//context.Set(req, "decoded", token.Claims)
					endpoint(c)
				} else {
					_ = json.NewEncoder(c.Writer).Encode(restwebapi.ErrorMsg{Message: "Invalid authorization token"})
				}
			} else {
				_ = json.NewEncoder(c.Writer).Encode(restwebapi.ErrorMsg{Message: "Invalid authorization token"})
			}
		} else {
			_ = json.NewEncoder(c.Writer).Encode(restwebapi.ErrorMsg{Message: "An authorization header is required"})
		}
	}
}
/*
func isAuthorizedOld(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		if c.Request.Header["Authorization"] != nil {

			bearerToken := c.Request.Header.Get("Authorization")
			jwtString := strings.Split(bearerToken, " ")[1]

			token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
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
*/
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
