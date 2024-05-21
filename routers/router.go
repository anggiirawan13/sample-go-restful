package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sample-go-restful/controllers"
)

func SetupRoutes(app *fiber.App, userController *controllers.UserController) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.Get("/", userController.GetAllUsers)
	users.Get("/:id", userController.GetUserByID)
	users.Post("/", userController.CreateUser)
	users.Put("/:id", userController.UpdateUser)
	users.Delete("/:id", userController.DeleteUser)
}
