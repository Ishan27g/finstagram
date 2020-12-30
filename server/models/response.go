package models

type HTTPResponse struct {
	Code     int
	Message  string
	Response interface{}
}
