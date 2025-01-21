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
	r.HandleFunc("/login", middlewares.ChequeoBD(routers.Login)).Methods("POST")
	r.HandleFunc("/verperfil", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.VerPerfil))).Methods("GET")
	r.HandleFunc("/modificarPerfil", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")
	r.HandleFunc("/graboTweet", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.GraboTweet))).Methods("POST")
	r.HandleFunc("/leoTweets", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.LeoTweets))).Methods("GET")
	r.HandleFunc("/borrarTweet", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.BorrarTweet))).Methods("DELETE")

	r.HandleFunc("/subirAvatar", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.SubirAvatar))).Methods("POST")
	r.HandleFunc("/obtenerAvatar", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.ObtenerAvatar))).Methods("GET")
	r.HandleFunc("/subirBanner", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.SubirBanner))).Methods("POST")
	r.HandleFunc("/obtenerBanner", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.ObtenerBanner))).Methods("GET")

	r.HandleFunc("/altaRelacion", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.AltaRelacion))).Methods("POST")
	r.HandleFunc("/bajaRelacion", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.BajaRelacion))).Methods("DELETE")
	r.HandleFunc("/consultaRelacion", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.ConsultaRelacion))).Methods("GET")

	r.HandleFunc("/listarUsuarios", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.ListaUsuarios))).Methods("GET")

	r.HandleFunc("/leoTweetsRelacion", middlewares.ChequeoBD(middlewares.ValidarJWT(routers.LeoTweetsRelacion))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Println("server corriendo en el puerto " + PORT)

	handler := cors.AllowAll().Handler(r)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
