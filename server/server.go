package main

import (
	`app/handlers`
	`app/routes`
	`github.com/gofiber/fiber/v2`
)


func main() {
	handlers.ConnectDatabase()

	application := fiber.New()

	// All valid api endpoints, route all /app to authentication middleware
	app := application.Group("/api")
	app.Use("/app",routes.Auth)

	routes.InitRoutes(app)

	_ = application.Listen(":3000")
}
