// Package authentication contains jwt signing and verification functions
package authentication

import (
    "errors"
    "fmt"
    "os"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("JWT_SECRET")

// GenerateJWTaccess creates an access token
func GenerateJWTaccess(email string, userID string) (string, error) {
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email":  email,
        "userId": userID,
        "exp":    time.Now().Add(time.Hour * 2).Unix(),
        "type":   "access",
    })
    return accessToken.SignedString([]byte(secretKey))
}

// GenerateJWTrefresh creates a refresh token
func GenerateJWTrefresh(email string, userID string) (string, error) {
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email":  email,
        "userId": userID,
        "exp":    time.Now().Add(time.Hour * 730).Unix(),
        "type":   "refresh",
    })
    return refreshToken.SignedString([]byte(secretKey))
}

// VerifyToken verifies the token
func VerifyToken(tokenString string) (string, string, string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if token.Method != jwt.SigningMethodHS256 {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secretKey), nil
    })


    if err != nil {
        return "", "","", fmt.Errorf("failed to parse token: %w", err)
    }

    if !token.Valid {
        return "", "", "", errors.New("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", "", "", errors.New("invalid claims format")
    }

    userID, ok := claims["userId"].(string)
    if !ok {
        return "", "","",  errors.New("invalid userID claim")
    }

    tokenType, ok := claims["type"].(string)
    if !ok {
        return "", "", "", errors.New("invalid token type claim")
    }

    email, ok := claims["email"].(string)
    if !ok{
        return "","","",  errors.New("invalid email claim")
  }

    return userID, tokenType, email, nil
}
