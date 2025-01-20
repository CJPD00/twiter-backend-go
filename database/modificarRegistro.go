package database

import (
	"context"
	"log"
	"time"

	"github.com/CJPD00/twiter-backend-go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModificoRegistro permite modificar un usuario en la base de datos, solo se modifica los campos que se envian en el json.
// Recibe un json con los datos del usuario a registrar, y el ID del usuario a modificar, y modifica los campos
// que se encuentren en el json en la base de datos. Retorna un booleano true si el registro es exitoso, de lo contrario
// se devuelve un error.
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-golang")

	col := db.Collection("usuarios")

	objID, _ := primitive.ObjectIDFromHex(ID)

	log.Printf("objID: %v", objID)

	condicion := bson.M{"_id": objID}

	usuario := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		usuario["nombre"] = u.Nombre
	}

	if len(u.Apellidos) > 0 {
		usuario["apellidos"] = u.Apellidos
	}

	if len(u.Avatar) > 0 {
		usuario["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		usuario["banner"] = u.Banner
	}

	if len(u.Biografia) > 0 {
		usuario["biografia"] = u.Biografia
	}

	if len(u.SitioWeb) > 0 {
		usuario["sitio_web"] = u.SitioWeb
	}

	usuario["fechaNacimiento"] = u.FechaNacimiento

	log.Printf("usuario a modificar: %v", usuario)

	update := bson.M{
		"$set": usuario,
	}

	_, err := col.UpdateOne(ctx, condicion, update)

	if err != nil {
		return false, err
	}

	return true, nil

}
