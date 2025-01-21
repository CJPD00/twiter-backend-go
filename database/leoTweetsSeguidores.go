package database

import (
	"context"
	"time"

	"github.com/CJPD00/twiter-backend-go/models"

	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsSeguidores(ID string, page int) ([]models.DevuelvoTweetsSeguidores, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-golang")
	col := db.Collection("relacion")

	skip := (page - 1) * 20

	condicion := make([]bson.M, 0)
	condicion = append(condicion, bson.M{"$match": bson.M{"usuarioid": ID}})

	condicion = append(condicion, bson.M{"$lookup": bson.M{"from": "tweet", "localField": "usuariorelacionid", "foreignField": "userid", "as": "tweet"}})
	condicion = append(condicion, bson.M{"$unwind": "$tweet"})
	condicion = append(condicion, bson.M{"$sort": bson.M{"fecha": -1}})

	condicion = append(condicion, bson.M{"$skip": skip})
	condicion = append(condicion, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, condicion)

	var result []models.DevuelvoTweetsSeguidores

	err := cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true

}
