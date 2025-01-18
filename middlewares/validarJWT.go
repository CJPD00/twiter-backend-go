package middlewares

import (
	"github.com/CJPD00/twiter-backend-go/routers"
	"net/http"
)

// ValidarJWT es un middleware que permite verificar el token de un usuario
// para asi poder acceder a los endpoints, si el token es invalido, se
// devuelve un error 400.
func ValidarJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
