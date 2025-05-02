package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	Role        string `json:"role"`
	jwt.StandardClaims
}
