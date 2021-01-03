package handlers

import (
	`app/models`
	`context`
	`github.com/gofiber/fiber/v2`
	`go.mongodb.org/mongo-driver/bson`
	`log`
	`time`
)

func GetUsers(c *fiber.Ctx) error {
	// Find all in database
	collection := Client.Database("DB-1").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	defer cancel()
	if err != nil {
		log.Println(err)
		return c.JSON(respond(401, "Not Found", nil))
	}
	var users []models.UserResponse
	for cur.Next(context.TODO()) {
		var user models.UserResponse
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if users == nil {
		return c.JSON(respond(401,"Not Found", nil))
	} else {
		return c.JSON(respond(200,"OK", users))
	}
}
func getUser(userID int) *models.UserResponse {
	collection := Client.Database("DB-1").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	cur, err := collection.Find(ctx, bson.D{{"userId", userID}})
	defer cancel()
	if err != nil {
		log.Println(err)
		return nil
	}
	var user models.User
	var userRsp models.UserResponse
	cur.Next(context.TODO())
	err = cur.Decode(&user)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	userRsp.Username = user.Username
	userRsp.Avatar = user.Avatar
	return &userRsp
}
