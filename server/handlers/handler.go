package handlers

import (
	`app/models`
	`context`
	`encoding/json`
	`github.com/gofiber/fiber/v2`
	`go.mongodb.org/mongo-driver/bson`
	`time`
)
func respond(code int, msg string) []byte{
	httpResponse := new(models.HTTPResponse)
	httpResponse.Code = code
	httpResponse.Message = msg
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	return jsonResponse
}

func AppRoot(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Signup(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil{
		return err
	}
	if user.Username == "" {
		return c.Send(respond(401,"Username can't be empty"))
	}
	if user.Password == "" {
		return c.Send(respond(401,"Username can't be empty"))
	}

	// Insert into database
	collection := Client.Database("DB-1").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, databaseErr := collection.InsertOne(ctx, bson.M{
		"password": user.Password,
		"username": user.Username,
	})
	defer cancel()
	if databaseErr != nil {
		return c.Send(respond(500,"Database error"))
	}

	return c.Send(respond(200,"Signed up"))
}
func Login(c *fiber.Ctx) error {
	userReq := new(models.User)
	err := c.BodyParser(userReq)
	if err != nil{
		return err
	}
	if userReq.Username == "" {
		return c.Send(respond(401,"Username can't be empty"))
	}
	if userReq.Password == "" {
		return c.Send(respond(401,"Username can't be empty"))
	}

	// Insert into database
	user:= new(models.User)
	collection := Client.Database("test").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = collection.FindOne(ctx, bson.M{
		"user":    userReq.Username,
		"password": userReq.Password,
	}).Decode(&user)

	defer cancel()
	if user == nil{
		return c.Send(respond(401,"Not Found"))
	} else {
		return c.Send(respond(200,"Logged in"))}
}