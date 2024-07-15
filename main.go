package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-test/routes"
)

func main() {
	app := fiber.New()

	routes.ApiRoutes(app) // for using in routes_api.go

	app.Listen(":3000")
}