package authentication

import (
	"context"
	"log"
	"net/http"

	"github.com/ank809/Event-Management-golang~/database"
	"github.com/ank809/Event-Management-golang~/helpers"
	"github.com/ank809/Event-Management-golang~/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user models.User

	user.ID = primitive.NewObjectID()
	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "Error in binding JSON")
		return
	}
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, "Name cannot be empty")
		return
	}
	isValidPassword, res := helpers.CheckPassword(user.Password)
	if !isValidPassword {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user.Password = string(hashedPassword)
	isValidEmail, res := helpers.VerifyEmail(user.Email)
	if !isValidEmail {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	isPhoneNumberValid, res := helpers.VerifyMobileNumber(user.PhoneNumber)
	if !isPhoneNumberValid {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if user.Role == "Attendee" {
		collection_name := "Attendee"
		coll := database.OpenCollection(database.Client, collection_name)
		res, err := coll.InsertOne(context.TODO(), user)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(res)
		c.JSON(http.StatusOK, "Attendee inserted successfully")

	} else {
		collection_name := "EventManager"
		coll := database.OpenCollection(database.Client, collection_name)
		res, err := coll.InsertOne(context.TODO(), user)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(res)
		c.JSON(http.StatusOK, "EventManager inserted successfully")
	}
}
