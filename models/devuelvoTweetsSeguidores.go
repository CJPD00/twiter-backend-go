package models

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// DevuelvoTweetsSeguidores es el modelo de tweet
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID            string             `bson:"usuarioid" json:"usuarioid,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"usuariorelacionid,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"id,omitempty"`
	}
}
