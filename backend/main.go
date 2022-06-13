package main

import (
	"backend/configs"
	"backend/routes" //add this
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app) //add this
PORT :=os.Getenv("PORT")
if PORT == " "{
	PORT = "8080"
}
	app.Listen(":"+PORT)
}
