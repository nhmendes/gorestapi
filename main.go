package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nhmendes/restapi/restwebapi"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySigningKey = []byte("mykey")

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "not authorized")
		}
	})
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter().StrictSlash(true)

	// Route handles & endpoints
	r.Handle("/books", isAuthorized(restwebapi.GetBooks)).Methods("GET")
	r.Handle("/books/{id}", isAuthorized(restwebapi.GetBook)).Methods("GET")
	r.Handle("/books", isAuthorized(restwebapi.CreateBook)).Methods("POST")
	r.Handle("/books/{id}", isAuthorized(restwebapi.UpdateBook)).Methods("PUT")
	r.Handle("/books/{id}", isAuthorized(restwebapi.DeleteBook)).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8001", r))
}

// Request sample
// {
// 	"isbn":"4545454",
// 	"title":"Book Three",
// 	"author":{"firstname":"Harry","lastname":"White"}
// }
