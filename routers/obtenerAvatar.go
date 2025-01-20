package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/CJPD00/twiter-backend-go/database"
)

// ObtenerAvatar devuelve la imagen avatar de un usuario por su ID
// Recibe el ID del usuario por parametro en la url, y devuelve el avatar
// que se encuentra en la carpeta uploads/avatar/
// Si el ID no existe, se devuelve un error de 400.
// Si el avatar no existe, se devuelve un error de 400.
// Si hubo un error al leer el archivo, se devuelve un error de 400.
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := database.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatar/" + perfil.Avatar)

	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
		return
	}

}
