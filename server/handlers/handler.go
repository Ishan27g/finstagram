package handlers

import (
	`app/models`
	`context`
	`fmt`
	`github.com/gofiber/fiber/v2`
	`github.com/google/uuid`
	`go.mongodb.org/mongo-driver/bson`
	`log`
	`time`
)
func respond(code int, msg string, msg2 interface{}) *models.HTTPResponse{
	httpResponse := new(models.HTTPResponse)
	httpResponse.Code = code
	httpResponse.Message = msg
	httpResponse.Response = msg2
	/*jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	log.Print("Response is: ")
	fmt.Println(jsonResponse)
	log.Print("")*/
	return httpResponse
}
func Signup(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil{
		return err
	}
	if user.Username == "" {
		return c.JSON(respond(401,"Username can't be empty", nil))
	}
	if user.Password == "" {
		return c.JSON(respond(401,"Password can't be empty", nil))
	}
	if user.Avatar == "" {
		return c.JSON(respond(401,"Avatar can't be empty", nil))
	}

	var posts []int
	uid := uuid.New().ID()
	// Insert into database
	collection := Client.Database("DB-1").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	_, databaseErr := collection.InsertOne(ctx, bson.M{
		"userId": uid,
		"password": user.Password,
		"username": user.Username,
		"email": user.Email,
		"posts": posts,
		"avatar": user.Avatar,
	})
	defer cancel()
	if databaseErr != nil {
		return c.JSON(respond(500,"Database error", nil))
	}
	fmt.Println("signed up")
	return c.JSON(respond(200,"Signed up", nil))
}
func Login(c *fiber.Ctx) error {
	userReq := new(models.User)
	err := c.BodyParser(userReq)
	if err != nil{
		return err
	}
	if userReq.Username == "" {
		return c.JSON(respond(401,"Username can't be empty", nil))
	}
	if userReq.Password == "" {
		return c.JSON(respond(401,"Username can't be empty", nil))
	}

	log.Print(userReq)
	// Find match in database
	user:= new(models.User)
	collection := Client.Database("DB-1").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	err = collection.FindOne(ctx, bson.M{
		"username": userReq.Username,
		"password": userReq.Password,
	}).Decode(&user)
	defer cancel()
	if err != nil {
		log.Println(err)
		return c.JSON(respond(401, "Not Found", nil))
	}
	if user == nil {
		return c.JSON(respond(401,"Not Found", nil))
	} else {
		return c.JSON(respond(200,"Logged in", user.UserId))
	}
}