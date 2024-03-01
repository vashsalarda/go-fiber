package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Golang API using Fiber & mongoDB"})
	})

	configs.ConnectDB()

	routes.UserRoute(app)

	// List all routes
	fmt.Println("API Routes:")

	routes := app.GetRoutes()
	for _, route := range routes {
		method := route.Method
		path := route.Path
		if path != "/" && method != "HEAD" {
			fmt.Printf("%-6s %-20s\n", method, path)
		}
	}

	app.Listen(":6000")
}
