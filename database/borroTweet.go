package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BorroTweet elimina un tweet en la base de datos, solo puede eliminar un tweet
// si el UserID coincide con el del tweet a eliminar.
// Recibe el ID del tweet a eliminar, y el UserID del usuario que lo va a eliminar.
// Retorna un booleano true si el registro es exitoso, de lo contrario
// se devuelve un error.
func BorroTweet(ID string, UserID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-golang")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condicion)

	if err != nil {
		return false, err
	}

	return true, nil
}
