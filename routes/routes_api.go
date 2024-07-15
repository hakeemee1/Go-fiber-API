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

	api := app.Group("/api")
	v1 := api.Group("/v1")
	fact := v1.Group("/fact")

	v1.Get("/", controllers.HelloTest)

	fact.Get("/:num", controllers.FactCalc)
	
}
