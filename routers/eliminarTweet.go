package routers

import (
	"github.com/CJPD00/twiter-backend-go/database"
	"net/http"
)

// BorrarTweet permite borrar un tweet en la base de datos, solo se puede borrar si el tweet es del usuario autenticado.
// Recibe el ID del tweet a eliminar, y devuelve un booleano true si el registro es exitoso, de lo contrario
// se devuelve un error.
func BorrarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	_, err := database.BorroTweet(ID, IDUsuario)

	if err != nil {
		http.Error(w, "Error al borrar el tweet "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
