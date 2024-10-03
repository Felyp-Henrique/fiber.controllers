package controllers

import (
	"Felyp-Henrique/fiber.controllers/controllers"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type WelcomeController struct {
	controllers.DefaultController
}

func (WelcomeController) Get(c *fiber.Ctx) error {
	return c.SendString(strings.Join([]string{
		"Welcome to Go Fiber Controllers! :D\n",
		"- Try call GET /users",
		"- Try call GET /users/:id",
		"- Try call POST /users",
	}, "\n"))
}
