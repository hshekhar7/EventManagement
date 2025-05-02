package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `json:"name"`
	DateTime       string             `json:"datetime"`
	Venue          string             `json:"venue"`
	TotalAttendees int                `json:"totalattendees"`
	Organizer      string             `json:"organizer"`
	Attendees      []RegisterUser     `json:"attendees"`
}
