package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	_ "github.com/mattn/go-sqlite3"

    "github.com/bagasjs/go-backend-template/app"
    "github.com/bagasjs/go-backend-template/http/controllers"
)

func Serve() {
    config := application.ApplicationConfig {}
    var app application.Application
    if err := app.Init(config); err != nil {
        log.Fatal(err)
    }
    defer app.Destroy()

    engine := html.New("./res/views/", ".html")
    rt := fiber.New(fiber.Config {
        Views: engine,
    })
    rt.Static("/public", "./res/public/")

    rt.Get("/admin", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to admin page")
    })

    // Define controllers here
    apiUserController := controllers.NewAPIUserController(app.UserService)
    apiJWTAuthController := controllers.NewAPIJWTAuthController(app.UserService)

    // Register controller to router here
    rt.Route("/api/users", apiUserController.Route)
    rt.Route("/api/auth", apiJWTAuthController.Route)

    log.Fatal(rt.Listen(":8080"))
}
