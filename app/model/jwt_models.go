package model

import "github.com/golang-jwt/jwt/v5"

type JWTAuthClaim struct {
    jwt.RegisteredClaims
    User GeneralUserResponse `json:"user"`
}
