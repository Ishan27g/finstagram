package handlers

import (
	`app/models`
	`github.com/gofiber/fiber/v2`
	`go.mongodb.org/mongo-driver/bson`
	`log`
	`strconv`
)


func GetPosts(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if userId == "" {
		return c.JSON(respond(401,"user's Id can't be empty", nil))
	}
	collectionPosts  :=Client.Database("DB-1").Collection("posts")
	posts, err := FindPostsInDB(collectionPosts, bson.D{})
	if err != nil {
		log.Println(err)
		return c.JSON(respond(401, "Not Found", nil))
	}
	if posts == nil{
		return c.JSON(respond(401,"Error finding posts", nil))
	} else {
		return c.JSON(respond(200,"Success", posts))
	}
}

func GetUserPosts(c *fiber.Ctx) error {
	userId := c.Params("userId")
	targetId := c.Params("targetId")
	if userId == "" || targetId == "" {
		return c.JSON(respond(401,"user's Id can't be empty", nil))
	}
	filter := bson.D{{"username", targetId}}
	collectionPosts  :=Client.Database("DB-1").Collection("posts")
	posts, err := FindPostsInDB(collectionPosts,filter)
	if err != nil {
		log.Println(err)
		return c.JSON(respond(401, "Not Found", nil))
	}
	if posts == nil{
		return c.JSON(respond(401,"Error finding posts", nil))
	} else {
		return c.JSON(respond(200,"Success", posts))
	}
}

func AddPosts(c *fiber.Ctx) error {
	var userId = c.Params("userId")
	id, err:= strconv.Atoi(userId)
	user:= getUser(id)
	if user == nil {
		return c.JSON(respond(401,"Error finding user", nil))
	}
	post := new(models.Post)
	err = c.BodyParser(post)
	if err != nil{
		return err
	}
	if userId == "" {
		return c.JSON(respond(401,"user's Id can't be empty", nil))
	}

	collectionPosts := Client.Database("DB-1").Collection("posts")
	databaseErr := AddPostToDB(collectionPosts, post, userId, user)
	if databaseErr != nil {
		return c.JSON(respond(500,"Database error", nil))
	}
	return c.JSON(respond(200,"Added post", nil))
}
