package authentication

import (
	"context"
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

func EventManagerLogin(c *gin.Context) {
	var manager models.User

	var foundManager models.User

	if err := c.BindJSON(&manager); err != nil {
		c.JSON(http.StatusBadRequest, "Error in binding json")
		return
	}

	collection_name := "EventManager"
	coll := database.OpenCollection(database.Client, collection_name)
	err := coll.FindOne(context.TODO(), bson.M{"name": manager.Name}).Decode(&foundManager)
	if err != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundManager.Password), []byte(manager.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Incorrect Password")
		return
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
		return
	}
	key := []byte(os.Getenv("JWT_KEY"))
	expiration_time := time.Now().Add(time.Minute * 10)
	claims := &models.Claims{
		Name:        foundManager.Name,
		Role:        foundManager.Role,
		PhoneNumber: foundManager.PhoneNumber,
		Email:       foundManager.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration_time.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(key)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "token",
		Value: tokenstring,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Manager logged in successfully",
		"token":   tokenstring,
	})
}
