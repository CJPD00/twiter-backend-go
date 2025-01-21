package database

import (
	"context"
	"log"
	"time"

	"github.com/CJPD00/twiter-backend-go/models"
	"go.mongodb.org/mongo-driver/bson"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]models.Usuario, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter-golang")
	col := db.Collection("usuarios")

	var resultado []models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)

	if err != nil {
		log.Println(err.Error())
		return resultado, false
	}

	var encontrado, incluir bool

	for cursor.Next(ctx) {

		var s models.Usuario
		err := cursor.Decode(&s)
		if err != nil {
			log.Println(err.Error())
			return resultado, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, _ = ConsultoRelacion(r)

		log.Println(encontrado)

		if tipo == "new" && !encontrado {
			incluir = true
		}

		if tipo == "follow" && encontrado {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
			resultado = append(resultado, s)
		}

	}

	err = cursor.Close(ctx)
	if err != nil {
		log.Println(err.Error())
		return resultado, false
	}

	return resultado, true

}
