package routes

import (
	`app/handlers`
	`github.com/gofiber/fiber/v2`
)

func InitRoutes(app fiber.Router){

	app.Post("/signup", handlers.Signup)
	app.Post("/login", handlers.Login)

	app.Get("/posts/:userId", handlers.GetPosts)
	app.Get("/posts/:userId/:targetId", handlers.GetUserPosts)
	app.Post("/posts/:userId", handlers.AddPosts)

	app.Get("/users", handlers.GetUsers)



}
