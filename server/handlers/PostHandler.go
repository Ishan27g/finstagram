package handlers

import (
	`app/models`
	`context`
	`fmt`
	`github.com/gofiber/fiber/v2`
	`go.mongodb.org/mongo-driver/bson`
	`log`
	`time`
)

func GetPosts(c *fiber.Ctx) error {
	log.Print("ree")
	userId := c.Params("userId")
	if userId == "" {
		return c.Send(respond(401,"user's Id can't be empty", nil))
	}
	log.Print(userId)
	collection := Client.Database("DB-1").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var posts []models.Post
	cur, err := collection.Find(ctx, bson.D{})
	defer cancel()
	if err != nil {
		log.Println(err)
		return c.Send(respond(401, "Not Found", nil))
	}
	for cur.Next(context.TODO()) {
		var post models.Post
		err := cur.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	if posts == nil{
		return c.Send(respond(401,"Error finding posts", nil))
	} else {
		return c.Send(respond(200,"Success", posts))
	}
}

func AddPosts(c *fiber.Ctx) error {
	userId := c.Params("userId")
	post := new(models.Post)
	err := c.BodyParser(post)
	if err != nil{
		return err
	}
	if userId == "" {
		return c.Send(respond(401,"user's Id can't be empty", nil))
	}
	log.Print(post)
	collection := Client.Database("DB-1").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, databaseErr := collection.InsertOne(ctx, bson.M{
		"caption": post.Caption,
		"imageUrl": post.ImagUrl,
		"date": post.Date,
	})
	defer cancel()
	if databaseErr != nil {
		return c.Send(respond(500,"Database error", nil))
	}
	fmt.Println("added post")
	return c.Send(respond(200,"Added post", nil))
}