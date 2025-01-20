package database

import (
	"context"
	"log"
	"time"

	"github.com/CJPD00/twiter-backend-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LeeTweets lee los tweets de un usuario en especifico
// Recibe el ID del usuario y la pagina a leer, y devuelve los tweets
// de ese usuario, y un booleano indicando si se encontraron o no
// tweets en la base de datos.
func LeoTweets(ID string, page int) ([]*models.DevuelvoTweet, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-golang")
	col := db.Collection("tweet")

	var resultado []*models.DevuelvoTweet

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((int64(page) - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultado, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.DevuelvoTweet

		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, false
		}

		resultado = append(resultado, &registro)
	}

	return resultado, true
}
