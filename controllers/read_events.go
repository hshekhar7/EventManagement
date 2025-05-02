package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/ank809/Event-Management-golang~/database"
	"github.com/ank809/Event-Management-golang~/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func ReadAllEvents(c *gin.Context) {
	var events []models.Event

	collection_name := "Events"
	coll := database.OpenCollection(database.Client, collection_name)
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Hello")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := cursor.All(context.TODO(), &events); err != nil {
		log.Println("Error decoding documents:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode documents"})
		return
	}
	if len(events) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "No events found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Documents": events,
	})

}
