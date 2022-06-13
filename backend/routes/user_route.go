package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/register",controllers.Register)
	// app.Post("/login",controllers.Login)
	app.Post("/product", controllers.CreateProduct)
	app.Get("/product/:productId", controllers.GetAProduct)
	app.Put("/product/:productId", controllers.EditAProduct)
	app.Delete("/product/:productId", controllers.DeleteAProduct)
	app.Get("/products", controllers.GetAllProducts)
}
