package routes

import (
	"Proj-GO/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App){

	app.Get("/", controllers.Index)
	app.Get("/show", controllers.Show)
	app.Get("/new", controllers.New)
	app.Get("/edit", controllers.Edit)

	app.Post("/insert", controllers.Insert)
	app.Get("/update", controllers.Update)
	app.Get("/delete", controllers.Delete)

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)
}
