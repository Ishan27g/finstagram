package models

// ErrorResponse is struct for sending error message with code.
type ErrorResponse struct {
	Code    int
	Message string
}

// SuccessResponse is struct for sending error message with code.
type SuccessResponse struct {
	Code     int
	Message  string
	Response interface{}
}

// RegistationParams is struct to read the request body
type RegistationParams struct {
	Username     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
// Signup
type Signup struct {
	Username     string `json:"username"`
	Password string `json:"password"`
}

// LoginParams is struct to read the request body
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SuccessfulLoginResponse is struct to send the request response
type SuccessfulLoginResponse struct {
	Email     string
	AuthToken string
}

// UserDetails is struct used for user details
type UserDetails struct {
	Username     string
	Email    string
	Password string
}
