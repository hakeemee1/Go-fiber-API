package routes

import (
	c "go-fiber-test/controllers"

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
	v3 := api.Group("/v3")
	fact := v1.Group("/fact")
	myName := v3.Group("/kimi")

	v1.Get("/", c.HelloTest)

	fact.Get("/:num", c.FactCalc)

	myName.Get("/", c.AsciiCalc)
	v1.Post("/register", c.Register)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)

}
