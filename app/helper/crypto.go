package helper

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyPasswordHash(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateToken(email string) (string, error) {
	saltSecret := GetEnv("SALT_TOKEN_SECRET", "SALTTOKENSECRET123")
	jwtKey := []byte(saltSecret)
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func GenerateRefresh(email string) (string, error) {
	saltSecret := GetEnv("SALT_REFRESH_SECRET", "SALTREFRESHSECRET123")
	jwtKey := []byte(saltSecret)
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*jwt.RegisteredClaims, error) {
	saltSecret := GetEnv("SALT_TOKEN_SECRET", "SALTTOKENSECRET123")
	jwtKey := []byte(saltSecret)

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func ValidateRefresh(tokenString string) (*jwt.RegisteredClaims, error) {
	saltSecret := GetEnv("SALT_REFRESH_SECRET", "SALTREFRESHSECRET123")
	jwtKey := []byte(saltSecret)

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func GetAuthHeader(r *http.Request) (string, error) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header missing")
	}

	// Extract the token from the Authorization header (format: Bearer <token>)
	tokenString := authHeader[len("Bearer "):]
	return tokenString, nil
}
