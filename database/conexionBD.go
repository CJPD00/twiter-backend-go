package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es el objeto de conexion a la BD
var MongoCN = ConectarBD()

// ConectarBD establece la conexion a la base de datos de mongoDB
func ConectarBD() *mongo.Client {
	godotenv.Load()
	var MONGO_URI = os.Getenv("MONGO_URI")
	var clientOptions = options.Client().ApplyURI(MONGO_URI)
	log.Println("Conectandose a la BD " + MONGO_URI)
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
