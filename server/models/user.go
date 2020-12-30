package models

type User struct {
	Id int
	Username string
	Password string
	Posts[] Post
	//Value int
}