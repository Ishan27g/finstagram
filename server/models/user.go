package models

type User struct {
	Id uint32
	Username string
	Password string
	Email string
	Posts []int
	//Value int
}