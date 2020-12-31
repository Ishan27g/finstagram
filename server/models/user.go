package models

type User struct {
	UserId uint32
	Username string
	Avatar string
	Password string
	Email string
	Posts []int
}

type UserResponse struct {
	Username string
	Avatar string
	Posts []int
	UserId uint32
}
