package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/CJPD00/twiter-backend-go/middlewares"
	"github.com/CJPD00/twiter-backend-go/routers"
)

// Manejadores inicia el servidor en el puerto definido en la variable PORT.
// La variable PORT debe ser un string con el numero de puerto al que se va
// a levantar el servidor. Si no se encuentra la variable, se levanta en el
// puerto 8080. El servidor utiliza las CORS para permitir peticiones desde
// cualquier dominio.
func Manejadores() {
	r := mux.NewRouter()

	r.HandleFunc("/registro", middlewares.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(r)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
