package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func ApiRoutes(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")
	profile := v1.Group("/profile")
	profile.Get("", c.GetUserProfiles)
	profile.Get("/ages", c.GetProfileAnyAges)
	profile.Get("/user", c.SearchProfiles)
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023"},
	}))
	//CRUD UserProfile
	profile.Post("/", c.AddUserProfile)
	profile.Put("/:id", c.UpdateUserProfile)
	profile.Delete("/:id", c.RemoveUserProfile)

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
	dog.Get("/idmorethan", c.GetDogsIdMoreThan)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)
	dog.Get("/remove", c.GetDeletedDogs)
	dog.Get("/dogssum", c.GetDogsJsonSummary)

	//Crud company
	company := v1.Group("/company")
	company.Get("", c.GetCompanies)
	company.Post("/", c.AddCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)

}
