package database

import (
	"context"
	"time"

	"github.com/CJPD00/twiter-backend-go/models"

	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario verifica si un usuario con el email dado ya existe en la base de datos.
// Retorna el modelo de usuario encontrado, un booleano indicando si el usuario fue encontrado o no,
// y el ID del usuario en formato string. Si no se encuentra el usuario, el booleano será false.
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-golang")

	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	} 

	return resultado, true, ID

}
