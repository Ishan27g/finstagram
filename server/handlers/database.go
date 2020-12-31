package handlers

import (
	`app/models`
	"context"
	`github.com/google/uuid`

	`go.mongodb.org/mongo-driver/bson`
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	`time`
)

var Client *mongo.Client

func ConnectDatabase() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	Client = client
	if err != nil {
		log.Fatal(err)
	}
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connected.")
}

func reversePosts(cur *mongo.Cursor) *[]models.Post{
	var posts []models.Post
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
	return &posts
}

func FindPostsInDB(collectionPosts *mongo.Collection, filter bson.D ) (*[]models.Post, error){
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	cur, err := collectionPosts.Find(ctx, filter)
	defer cancel()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	posts:= reversePosts(cur)
	return posts, nil
}

func AddPostToDB(collectionPosts *mongo.Collection, post *models.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	uid := uuid.New().ID()
	current := time.Now().UTC().Format("02 Jan 06 15:04 AEST")
	_, databaseErr := collectionPosts.InsertOne(ctx, bson.M{
		"id":       uid,
		"caption":  post.Caption,
		"imageUrl": post.ImageUrl,
		"date":     current,
		"username": post.Username,
	})
	if databaseErr != nil {
		return databaseErr
	}
	return nil
}

