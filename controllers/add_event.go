package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ank809/Event-Management-golang~/database"
	"github.com/ank809/Event-Management-golang~/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddEvent(c *gin.Context) {
	var event models.Event

	if err := c.BindJSON(&event); err != nil {
		log.Println("Error in binding Json")
		return
	}
	if event.Name == "" {
		c.JSON(http.StatusBadRequest, "Event name can't be empty")
		return
	}
	if event.Venue == "" {
		c.JSON(http.StatusBadRequest, "Event should have a venue")
		return
	}
	event.ID = primitive.NewObjectID()
	event.TotalAttendees = 0
	event.Attendees = []models.RegisterUser{}
	parsedTime, err := time.Parse(time.RFC1123, event.DateTime)
	if err != nil {
		log.Println("Error parsing date time:", err)
		c.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}
	event.DateTime = parsedTime.String()
	collection_name := "Events"
	coll := database.OpenCollection(database.Client, collection_name)
	res, err := coll.InsertOne(context.Background(), event)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "Event Added Successfully",
		"messgae": res,
	})

}
