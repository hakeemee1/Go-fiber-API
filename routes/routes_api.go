package routes

import (
	"go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func ApiRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566"},
	}))


	app.Get("/", controllers.HelloTest)
	//group api
	// api := app.Group("/api/v1")
	// fact := api.Group("/fact")
}
