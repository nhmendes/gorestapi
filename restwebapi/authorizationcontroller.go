package restwebapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// User ...
// Custom object which can be stored in the claims
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthToken ...
// This is what is retured to the user
type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// AuthTokenClaim ...
// This is the cliam object which gets parsed from the authorization header
type AuthTokenClaim struct {
	*jwt.StandardClaims
	User
}

// ErrorMsg ...
// Custom error object
type ErrorMsg struct {
	Message string `json:"message"`
}

// GetToken - Generates a new token (JWT)
func GetToken(c *gin.Context) {
	var user User
	_ = json.NewDecoder(c.Request.Body).Decode(&user)

	expiresAt := time.Now().Add(time.Hour * 1).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &AuthTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		User{user.Username, user.Password},
	}

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	c.Writer.WriteHeader(http.StatusOK)
	encodeError := json.NewEncoder(c.Writer).Encode(AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	})

	if encodeError != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

// GetTokenOld - Generates a new token (JWT)
func GetTokenOld(c *gin.Context) {

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
