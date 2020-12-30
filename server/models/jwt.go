package models

import `github.com/dgrijalva/jwt-go`

// Claims is  a struct that will be encoded to a JWT.
// jwt.StandardClaims is an embedded type to provide expiry time
type Claims struct {
	Email string
	jwt.StandardClaims
}
