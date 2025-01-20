package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/CJPD00/twiter-backend-go/database"
	"github.com/CJPD00/twiter-backend-go/models"
)

// SubirBanner sube la imagen banner del usuario
// Recibe un archivo y lo sube a la carpeta /uploads/avatar/
// con el nombre del ID del usuario y la extension del archivo
// que se sube.
// Devuelve un status 200 si la imagen se subio correctamente
// y un status 400 si hubo un error.
func SubirBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banner/" + IDUsuario + "." + extension

	if err != nil {
		http.Error(w, "Error al subir la imagen "+err.Error(), 400)
		return
	}

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen "+err.Error(), 400)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), 400)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = database.ModificoRegistro(usuario, IDUsuario)

	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Error al copiar la imagen ", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

}
