package routes

import (
	"Proj-GO/controllers"

	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App){

	app.Get("/", controllers.Index)
	app.Get("/show", controllers.Show)
	app.Get("/new", controllers.New)
	app.Get("/edit", controllers.Edit)

	app.Get("/insert", controllers.Insert)
	app.Get("/update", controllers.Update)
	app.Get("/delete", controllers.Delete)
}
