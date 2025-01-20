package models

// Tweet es el modelo de tweet
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
