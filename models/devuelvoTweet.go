package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DevuelvoTweet es el modelo de tweet
type DevuelvoTweet struct {
	ID      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
