package controllers

import (
	"net/http"

	"github.com/bagasjs/go-backend-template/app/model"
	"github.com/bagasjs/go-backend-template/app/service"
	"github.com/bagasjs/go-backend-template/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

type APIUserController struct {
    userService *service.UserService
}

func (controller *APIUserController) allUsers(c *fiber.Ctx) error {
    users, error := controller.userService.List()
    if error != nil {
        return c.JSON(fiber.Map {
            "message" : error.Message,
            "code" : error.Code,
            "data" : nil,
        })
    }

    return c.JSON(fiber.Map {
        "message" : "Users fetched",
        "code" : http.StatusOK,
        "data" : users,
    })
}

func (controller *APIUserController) viewUser(c *fiber.Ctx) error {
    userID, err := c.ParamsInt("id")
    if err != nil {
        return c.JSON(fiber.Map {
            "message" : "Parameter ID should type of integer",
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    user, error := controller.userService.Find(userID)
    if error != nil {
        return c.JSON(fiber.Map {
            "message" : error.Message,
            "code" : error.Code,
            "data" : nil,
        })
    }

    return c.JSON(user)
}

func (controller *APIUserController) storeUser(c *fiber.Ctx) error {
    request := model.CreateUpdateUserRequest{}
    if err := c.BodyParser(&request); err != nil {
        return c.JSON(fiber.Map{
            "message" : err.Error(),
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }
    err := controller.userService.Create(request)
    if err != nil {
        return c.JSON(fiber.Map {
            "message" : err.Message,
            "code" : err.Code,
            "data" : nil,
        })
    }

    return c.JSON(fiber.Map{
        "message" : "Creating user successful",
        "code" : http.StatusOK,
        "data" : nil,
    })
}

func (controller *APIUserController) updateUser(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return c.JSON(fiber.Map{
            "message" : err.Error(),
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    request := model.CreateUpdateUserRequest{}
    if err := c.BodyParser(&request); err != nil {
        return c.JSON(fiber.Map{
            "message" : err.Error(),
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    if err := controller.userService.Update(id, request); err != nil { 
        return c.JSON(fiber.Map {
            "message" : err.Message,
            "code" : err.Code,
            "data" : nil,
        })
    }

    return c.JSON(fiber.Map{
        "message" : "Updating user successful",
        "code" : http.StatusOK,
        "data" : nil,
    })
}

func (controller *APIUserController) destroyUser(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return c.JSON(fiber.Map{
            "message" : err.Error(),
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    if err := controller.userService.Destroy(id); err != nil { 
        return c.JSON(fiber.Map {
            "message" : err.Message,
            "code" : err.Code,
            "data" : nil,
        })
    }

    return c.JSON(fiber.Map{
        "message" : "Destroying user successful",
        "code" : http.StatusOK,
        "data" : nil,
    })
}

func (controller *APIUserController) Route(r fiber.Router) {
    r.Get("/", controller.allUsers)
    r.Post("/", controller.storeUser)
    r.Get("/:id", controller.viewUser)
    r.Put("/:id", middlewares.Authenticated, controller.updateUser)
    r.Delete("/:id", middlewares.Authenticated, controller.destroyUser)
}

func NewAPIUserController(userService *service.UserService) *APIUserController {
    return &APIUserController {
        userService: userService,
    }
}
