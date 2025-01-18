package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es el objeto de conexion a la BD
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/twitter-golang")

// ConectarBD establece la conexion a la base de datos de mongoDB
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la BD")
	return client
}

// ChequeoConnection permite saber si la conexion a la BD es exitosa
// retorna 1 si la conexion es exitosa, 0 en caso contrario
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
