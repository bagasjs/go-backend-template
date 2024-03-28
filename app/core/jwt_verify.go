package core

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWTString(tokenString string, signingKey string) error {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token,) (interface{}, error) {
        return []byte(signingKey), nil
    })

    if err != nil {
        return err
    }

    if !token.Valid {
        return fmt.Errorf("Invalid token")
    }

    return nil
}
