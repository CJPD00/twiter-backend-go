package database

import (
	"context"
	"time"

	"github.com/CJPD00/twiter-backend-go/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro permite insertar un nuevo usuario en la base de datos
// Recibe un json con los datos del usuario a registrar, verifica que
// el email no este duplicado, y que la contrasen ̃a tenga al menos 6
// caracteres. Si el registro es exitoso, se devuelve el ID del usuario
// en formato string y un booleano true, de lo contrario se devuelve un
// error.
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-golang")

	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
