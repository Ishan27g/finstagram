package routes

import (
	`fmt`
	`github.com/gofiber/fiber/v2`
)

func Auth(c *fiber.Ctx) error {
	//user:= new(models.User)
	fmt.Println("Auth Middleware hit !")
/*	err := c.BodyParser(user)
	if err != nil{
		log.Print(err)
		return err
	}
	log.Print(user.Username)
*/	return c.Next()
}
