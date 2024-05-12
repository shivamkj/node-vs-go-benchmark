package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("c7fefec7e3c0d46dd5da07ff172e7c61bd982a63071fcff9ba24165bc6753a66")

// GenerateToken generates a new JWT token with the given user ID
func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix() // Token expires in 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
