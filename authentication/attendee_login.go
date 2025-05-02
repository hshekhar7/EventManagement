package authentication

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ank809/Event-Management-golang~/database"
	"github.com/ank809/Event-Management-golang~/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func LoginAttendee(c *gin.Context) {
	var user models.User
	var foundUser models.User

	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		return
	}
	collection_name := "Attendee"
	coll := database.OpenCollection(database.Client, collection_name)
	err := coll.FindOne(context.TODO(), bson.M{"name": user.Name}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "User not found")
		log.Println(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Incorrrect password")
		return
	}

	expiration_time := time.Now().Add(time.Minute * 10)

	claims := &models.Claims{
		Name:        foundUser.Name,
		Role:        foundUser.Role,
		Email:       foundUser.Email,
		PhoneNumber: foundUser.PhoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration_time.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
		return
	}
	jwt_key := []byte(os.Getenv("JWT_KEY"))
	tokenString, err := token.SignedString(jwt_key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expiration_time,
	})
	c.JSON(http.StatusOK, gin.H{"token": tokenString,
		"success": "User login Successfully"})
}
