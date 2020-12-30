package handlers

import (
	`app/models`
	`context`
	`encoding/json`
	`fmt`
	`github.com/gofiber/fiber/v2`
	`github.com/google/uuid`
	`go.mongodb.org/mongo-driver/bson`
	`log`
	`time`
)
func respond(code int, msg string, msg2 interface{}) []byte{
	httpResponse := new(models.HTTPResponse)
	httpResponse.Code = code
	httpResponse.Message = msg
	httpResponse.Response = msg2
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
		return c.Send(respond(401,"Username can't be empty", nil))
	}
	if user.Password == "" {
		return c.Send(respond(401,"Username can't be empty", nil))
	}

	var posts []int
	uid := uuid.New().ID()
	// Insert into database
	collection := Client.Database("DB-1").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, databaseErr := collection.InsertOne(ctx, bson.M{
		"id": uid,
		"password": user.Password,
		"username": user.Username,
		"email": user.Email,
		"posts": posts,
	})
	defer cancel()
	if databaseErr != nil {
		return c.Send(respond(500,"Database error", nil))
	}
	fmt.Println("signed up")
	return c.Send(respond(200,"Signed up", nil))
}
func Login(c *fiber.Ctx) error {
	userReq := new(models.User)
	err := c.BodyParser(userReq)
	if err != nil{
		return err
	}
	if userReq.Username == "" {
		return c.Send(respond(401,"Username can't be empty", nil))
	}
	if userReq.Password == "" {
		return c.Send(respond(401,"Username can't be empty", nil))
	}

	log.Print(userReq)
	// Find match in database
	user:= new(models.User)
	collection := Client.Database("DB-1").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = collection.FindOne(ctx, bson.M{
		"username": userReq.Username,
		"password": userReq.Password,
	}).Decode(&user)
	defer cancel()
	if err != nil {
		log.Println(err)
		return c.Send(respond(401, "Not Found", nil))
	}
	if user == nil || err != nil{
		return c.Send(respond(401,"Not Found", nil))
	} else {
		return c.Send(respond(200,"Logged in", user.Id))
	}
}