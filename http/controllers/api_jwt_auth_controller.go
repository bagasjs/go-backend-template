package controllers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bagasjs/go-backend-template/app/core"
	"github.com/bagasjs/go-backend-template/app/model"
	"github.com/bagasjs/go-backend-template/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type APIJWTAuthController struct {
    userService *service.UserService
}

func NewAPIJWTAuthController(userService *service.UserService) *APIJWTAuthController {
    return &APIJWTAuthController{
        userService: userService,
    }
}

func (controller *APIJWTAuthController) login(c *fiber.Ctx) error {
    request := model.LoginUserRequest{}
    if err := c.BodyParser(&request); err != nil {
        return c.JSON(fiber.Map{
            "message" : err.Error(),
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    query := core.NewQueryBuilder().Where("email", "=", request.Email)
    users, err := controller.userService.UserRepository.Query(query)
    if err != nil {
        return c.JSON(fiber.Map{
            "message" : err.Message,
            "code" : err.Code,
            "data" : nil,
        })
    }
    if len(users) < 1 {
        return c.JSON(fiber.Map{
            "message" : "Invalid email or password",
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    user := users[0]

    if strings.Compare(user.Password, request.Password) != 0 {
        return c.JSON(fiber.Map{
            "message" : "Invalid email or password",
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    claims := &model.JWTAuthClaim{}
    claims.User.ID = user.ID
    claims.User.Name = user.Name
    claims.User.Email = user.Email
    claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 7))
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signingKey := os.Getenv("JWT_SIGNING_KEY")
    t, err2 := token.SignedString([]byte(signingKey))
    if err2 != nil {
        return c.JSON(fiber.Map{
            "message" : "Failed to sign JWT token",
            "code" : http.StatusInternalServerError,
            "data" : nil,
        })
    }

    return c.JSON(fiber.Map{
        "message" : "Login success. Retrieve your token!",
        "code" : http.StatusOK,
        "data" : t,
    })
}

func (controller *APIJWTAuthController) Route(r fiber.Router) {
    r.Post("/login", controller.login)
}
