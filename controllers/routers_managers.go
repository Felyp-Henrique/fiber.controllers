package controllers

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type RouterManager struct {
	App        *fiber.App
	Controller Controller
}

func (r *RouterManager) RegisterHandler(path string, controllerMethod reflect.Method, fiberCallback func(string, ...fiber.Handler) fiber.Router) {
	fmt.Println("Registering path", r.GetPathComplete(path))
	fiberCallback(r.GetPathComplete(path), func(c *fiber.Ctx) error {
		var (
			controllerReference          reflect.Value
			fiberContextReference        reflect.Value
			parameters                   []reflect.Value
			methodReturnsValues          []reflect.Value
			methodReturnValueAsInterface any
		)
		controllerReference = reflect.ValueOf(r.Controller)
		fiberContextReference = reflect.ValueOf(c)
		parameters = []reflect.Value{controllerReference, fiberContextReference}
		methodReturnsValues = controllerMethod.Func.Call(parameters)
		methodReturnValueAsInterface = methodReturnsValues[0].Interface()
		if methodReturnValueAsInterface != nil {
			return methodReturnValueAsInterface.(error)
		} else {
			return nil
		}
	})
}

func (r *RouterManager) RegisterHandlerGet(path string, controllerMethod reflect.Method) {
	r.RegisterHandler(path, controllerMethod, r.App.Get)
}

func (r *RouterManager) RegisterHandlerPost(path string, controllerMethod reflect.Method) {
	r.RegisterHandler(path, controllerMethod, r.App.Post)
}

func (r *RouterManager) RegisterHandlerPut(path string, controllerMethod reflect.Method) {
	r.RegisterHandler(path, controllerMethod, r.App.Put)
}

func (r *RouterManager) RegisterHandlerDelete(path string, controllerMethod reflect.Method) {
	r.RegisterHandler(path, controllerMethod, r.App.Delete)
}

func (r *RouterManager) GetPathComplete(path string) string {
	newPath, _ := url.JoinPath("/", r.Controller.Route(), path)
	return newPath
}

func NewRouterManager(app *fiber.App, controller Controller) *RouterManager {
	return &RouterManager{
		App:        app,
		Controller: controller,
	}
}
