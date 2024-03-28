package controllers

import "github.com/gofiber/fiber/v2"

type UserAdminController struct {

}

func (controller *UserAdminController) allUsers(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{})
}

func (controller *UserAdminController) viewUser(c *fiber.Ctx) error {
    return c.SendString("You are viewing an user")
}

func (controller *UserAdminController) storeUser(c *fiber.Ctx) error {
    return c.SendString("New user has been stored")
}

func (controller *UserAdminController) updateUser(c *fiber.Ctx) error {
    return c.SendString("New user has been created")
}

func (controller *UserAdminController) destroyUser(c *fiber.Ctx) error {
    return c.SendString("You are deleting an user")
}

func (controller *UserAdminController) Route(r fiber.Router) {
    r.Get("/", controller.allUsers)
    r.Post("/", controller.storeUser)
    r.Get("/:id", controller.viewUser)
    r.Put("/:id", controller.updateUser)
    r.Delete("/:id", controller.destroyUser)
}
