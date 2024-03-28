package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bagasjs/go-backend-template/app/core"
	"github.com/gofiber/fiber/v2"
)

const AUTHORIZATION_HEADER_PREFIX = "Bearer "
const AUTHORIZATION_HEADER_PREFIX_LENGTH = len(AUTHORIZATION_HEADER_PREFIX)

func Authenticated(c *fiber.Ctx) error {
    authorization := string(c.Request().Header.Peek("Authorization"))
    if len(authorization) < AUTHORIZATION_HEADER_PREFIX_LENGTH {
        return c.JSON(fiber.Map {
            "message" : "Authorization headers is not found or invalid",
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    tokenString := authorization[AUTHORIZATION_HEADER_PREFIX_LENGTH:]
    if len(tokenString) < 0 {
        return c.JSON(fiber.Map {
            "message" : "Authorization headers is not found or invalid",
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    if err := core.VerifyJWTString(tokenString, os.Getenv("JWT_SIGNING_KEY")); err != nil {
        return c.JSON(fiber.Map {
            "message" : fmt.Sprintf("Unauthenticated user: %s", err),
            "code" : http.StatusUnauthorized,
            "data" : nil,
        })
    }

    return c.Next()
}
