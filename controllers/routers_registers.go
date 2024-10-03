package controllers

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(app *fiber.App, controller Controller) {
	var (
		controllerType   reflect.Type
		controllerMethod reflect.Method
		routerParser     *RouterParser
		routerManager    *RouterManager
		routerPath       string
	)
	routerParser = NewRouterParser()
	routerManager = NewRouterManager(app, controller)
	controllerType = reflect.ValueOf(controller).Type()
	for index := range controllerType.NumMethod() {
		controllerMethod = controllerType.Method(index)
		routerPath = routerParser.GetPathComplete(controllerMethod)
		if routerParser.IsHandlerGetHttpMethod(controllerMethod) {
			routerManager.RegisterHandlerGet(routerPath, controllerMethod)
			continue
		}
		if routerParser.IsHandlerPostHttpMethod(controllerMethod) {
			routerManager.RegisterHandlerPost(routerPath, controllerMethod)
			continue
		}
		if routerParser.IsHandlerPutHttpMethod(controllerMethod) {
			routerManager.RegisterHandlerPut(routerPath, controllerMethod)
			continue
		}
		if routerParser.IsHandlerDeleteHttpMethod(controllerMethod) {
			routerManager.RegisterHandlerDelete(routerPath, controllerMethod)
			continue
		}
	}
}
