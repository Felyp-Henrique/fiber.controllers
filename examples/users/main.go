package main

import (
	"Felyp-Henrique/fiber.controllers/controllers"
	exampleController "Felyp-Henrique/fiber.controllers/examples/users/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	controllers.RegisterRouter(app, &exampleController.WelcomeController{})
	controllers.RegisterRouter(app, &exampleController.UsersController{})

	app.Listen(":8080")
}
