package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here
	app.Get("/users", controllers.GetAllUsers)
	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:id", controllers.GetAUser)
	app.Patch("/users/:id", controllers.EditAUser)
	app.Delete("/users/:id", controllers.DeleteAUser)
}
