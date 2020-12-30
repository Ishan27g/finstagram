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

	// All valid api endpoints, route all /app to authentication middleware
	app := application.Group("/api")
	app.Use(cors.New())
	routes.InitRoutes(app)

	_ = application.Listen(":3000")
}
