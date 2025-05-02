package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ank809/Event-Management-golang~/database"
	"github.com/ank809/Event-Management-golang~/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var newEvent models.Event
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err = c.BindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	collection_name := "Events"
	coll := database.OpenCollection(database.Client, collection_name)

	updatedEvent := bson.M{
		"$set": bson.M{
			"name":      newEvent.Name,
			"datetime":  newEvent.DateTime,
			"venue":     newEvent.Venue,
			"attendees": newEvent.Attendees,
			"organizer": newEvent.Organizer,
		},
	}
	var res bson.M
	err = coll.FindOneAndUpdate(context.TODO(), bson.M{"_id": objectId}, updatedEvent).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, "No documents found")
			return
		} else {
			c.JSON(http.StatusBadRequest, err)
			return
		}
	} else {
		message := fmt.Sprintf("Document updated successfully  %v", res)
		c.JSON(http.StatusOK, message)
	}

}
