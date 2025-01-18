package database

import (
	"context"
	"fmt"
	"time"

	"github.com/CJPD00/twiter-backend-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BuscoPerfil busca un perfil en la base de datos por su ID y devuelve el modelo de usuario encontrado.
// Si no se encuentra el perfil, se devuelve un error.
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-golang")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""

	if err != nil {
		fmt.Println("Registro no encontrado")
		return perfil, err
	}

	return perfil, nil

}
