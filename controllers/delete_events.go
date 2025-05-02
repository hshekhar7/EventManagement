package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ank809/Event-Management-golang~/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func DeleteEvents(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid object Id")
		return
	}
	collection_name := "Events"
	coll := database.OpenCollection(database.Client, collection_name)
	var res bson.M
	err = coll.FindOneAndDelete(context.TODO(), bson.M{"_id": objectId}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, "No document found")
		} else {
			c.JSON(http.StatusNotFound, err)
		}
	} else {
		message := fmt.Sprintf("Docuemtn deleted successfully %v", res)
		c.JSON(http.StatusOK, message)
	}

}
