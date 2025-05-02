package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/ank809/Event-Management-golang~/database"
	"github.com/ank809/Event-Management-golang~/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func RegisterEvent(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userClaims, exists := c.Get("user")
	log.Println("User claims", userClaims)
	if !exists {
		c.JSON(http.StatusInternalServerError, "User not found")
		return
	}
	claims := userClaims.(*models.Claims)
	log.Println("Current user", claims)
	var user models.RegisterUser = models.RegisterUser{
		ID:          primitive.NewObjectID(),
		Name:        claims.Name,
		Email:       claims.Email,
		PhoneNumber: claims.PhoneNumber,
	}

	collection_name := "Events"
	coll := database.OpenCollection(database.Client, collection_name)
	var res models.Event
	err = coll.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, "No documents found")
			return
		} else {
			c.JSON(http.StatusNotFound, err)
			return
		}
	} else {
		res.TotalAttendees += 1
		res.Attendees = append(res.Attendees, user)
		update := bson.M{
			"$set": bson.M{
				"totalattendees": res.TotalAttendees,
				"attendees":      res.Attendees,
			},
		}
		_, err := coll.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, "Registered successfully")

}
