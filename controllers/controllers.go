package controllers

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Route() string

	RouterCustom(*fiber.App)
}

type DefaultController struct {
}

func (DefaultController) Route() string {
	return "/"
}

func (DefaultController) RouterCustom(app *fiber.App) {

}
