package models

import "github.com/golang-jwt/jwt/v5"

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTClaim struct {
	Email            string               `json:"email"`
	RegisteredClaims jwt.RegisteredClaims `json:"registered_claims"`
}
