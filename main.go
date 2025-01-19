package main

import (
	"log"

	"github.com/CJPD00/twiter-backend-go/database"
	"github.com/CJPD00/twiter-backend-go/handlers"
	"github.com/joho/godotenv"
)

// main inicia el servidor en el puerto definido en la variable PORT.
// La variable PORT debe ser un string con el numero de puerto al que se va
// a levantar el servidor. Si no se encuentra la variable, se levanta en el
// puerto 8080. El servidor utiliza las CORS para permitir peticiones desde
// cualquier dominio. Chequea la conexion a la BD y si falla, termina la
// ejecucion del programa.
func main() {

	godotenv.Load()
	if database.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
	}
	handlers.Manejadores()

}
