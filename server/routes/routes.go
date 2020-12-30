package routes

import (
	`app/handlers`
	`github.com/gofiber/fiber/v2`
)

func InitRoutes(app fiber.Router){

	app.Get("/", handlers.AppRoot)

	app.Post("/signup", handlers.Signup)
	app.Post("/login", handlers.Login)


}