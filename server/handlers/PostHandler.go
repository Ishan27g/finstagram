package handlers

import (
	`app/models`
	`context`
	"fmt"
	`github.com/gofiber/fiber/v2`
	`github.com/google/uuid`
	`go.mongodb.org/mongo-driver/bson`
	`log`
	`strconv`
	`time`
)

func GetPosts(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if userId == "" {
		return c.JSON(respond(401,"user's Id can't be empty", nil))
	}
	collection := Client.Database("DB-1").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	var posts []models.Post
	cur, err := collection.Find(ctx, bson.D{})
	defer cancel()
	if err != nil {
		log.Println(err)
		return c.JSON(respond(401, "Not Found", nil))
	}
	for cur.Next(context.TODO()) {
		var post models.Post
		err := cur.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	for i := 0; i < len(posts)/2; i++ {
	j := len(posts) - i - 1
		posts[i], posts[j] = posts[j], posts[i]
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
	fmt.Print(userId, id)
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
	uid := uuid.New().ID()
	current := time.Now().UTC().Format("02 Jan 06 15:04 AEST")
	userid, _ := strconv.Atoi(userId)
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	collection := Client.Database("DB-1").Collection("posts")
	_, databaseErr := collection.InsertOne(ctx, bson.M{
		"id": uid,
		"caption": post.Caption,
		"imageUrl": post.ImageUrl,
		"date": current,
		"userId":userid,
		"user": user,
	})
	defer cancel()
	if databaseErr != nil {
		return c.JSON(respond(500,"Database error", nil))
	}
	return c.JSON(respond(200,"Added post", nil))
}
