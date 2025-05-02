package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name"`
	Email       string             `json:"email"`
	PhoneNumber string             `json:"phonenumber"`
	Role        string             `json:"role"`
	Password    string             `json:"password"`
}

type RegisterUser struct {
	Name        string             `json:"name"`
	ID          primitive.ObjectID `bson:"_id"`
	Email       string             `json:"email"`
	PhoneNumber string             `json:"phonenumber"`
}
