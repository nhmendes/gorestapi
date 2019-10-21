package restwebapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetToken - Generates a new token (JWT)
func GetToken(c *gin.Context) {

	mySigningKey := []byte("AllYourBase")
	/*
		// Create the Claims
		claims := &jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		}*/

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	/*token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})*/

	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.New(jwt.SigningMethodHS256)

	// Sign and get the complete encoded token as a string using the secret
	result, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf("failed to create new token, error: %v", err)
	}

	c.Writer.WriteHeader(http.StatusOK)
	encodeError := json.NewEncoder(c.Writer).Encode(result)

	if encodeError != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}
