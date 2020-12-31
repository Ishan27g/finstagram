package routes

import (
	`fmt`
	`github.com/gofiber/fiber/v2`
)

func Auth(c *fiber.Ctx) error {
	fmt.Print(c.IP() + "--->")
	fmt.Print(c.Method() + " ")
	fmt.Print(c.Path() + "\n")
	return c.Next()
}
