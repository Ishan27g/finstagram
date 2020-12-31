package main

import (
	`app/handlers`
	`app/routes`
	`github.com/gofiber/fiber/v2`
	`github.com/gofiber/fiber/v2/middleware/cors`
)

func main() {
	handlers.ConnectDatabase()

	application := fiber.New()

	app := application.Group("/api")
	app.Use(cors.New())
	app.Use("/",routes.Auth)
	routes.InitRoutes(app)

	_ = application.Listen(":3000")
}
