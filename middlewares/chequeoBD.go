package middlewares

import (
	"github.com/CJPD00/twiter-backend-go/database"
	"net/http"
)

// ChequeoBD es un middleware que permite chequear la conexion a la BD
// antes de que se ejecute una funcion. Si la conexion esta perdida,
// se devuelve un error 500. De lo contrario, se ejecuta la funcion dada.
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

