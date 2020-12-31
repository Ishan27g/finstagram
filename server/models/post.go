package models

type Post struct {
	Id uint32       `json:"Id"`
	Caption string  `json:"Caption"`
	ImageUrl string `json:"ImageUrl"`
	Date string     `json:"Date"`
	UserId uint32   `json:"UserId"`
}


