// Package auth jwt is an implementation of jwt authentication
package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// secretKey gotten from the env
var secretKey = os.Getenv("JWT_SECRET")

// GenerateJWTaccess access_token creates a token based on incoming request body
func GenerateJWTaccess(email string, userID string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
    "type":"access",
	})

	return accessToken.SignedString([]byte(secretKey))
}

// GenerateJWTrefresh refresh_token
func GenerateJWTrefresh(email string, userID string) (string, error){
  refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":  email,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 730).Unix(),
    "type": "refresh",
  })
  return refreshToken.SignedString([]byte(secretKey))
}

// VerifyToken verifies jwt token
func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("Could not parse tokens")
		return 0, errors.New("could not parse tokens")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid tokens")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid tokens claims")
	}

	userID := int64(claims["userID"].(float64))
	return userID, nil
} 
