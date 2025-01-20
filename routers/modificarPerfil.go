package routers

import (
	"encoding/json"
	"github.com/CJPD00/twiter-backend-go/database"
	"github.com/CJPD00/twiter-backend-go/helpers"
	"github.com/CJPD00/twiter-backend-go/models"
	"log"
	"net/http"
)

// ModificarPerfil permite modificar un usuario en la base de datos, solo se modifica los campos que se envian en el json.
// Recibe un json con los datos del usuario a registrar, y el ID del usuario a modificar, y modifica los campos
// que se encuentren en el json en la base de datos. Retorna un booleano true si el registro es exitoso, de lo contrario
// se devuelve un error.
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	log.Println("Registro recibido: ", t)

	if err != nil {
		http.Error(w, "Error en la peticion "+err.Error(), 400)
		return
	}

	if len(t.Email) > 0 {
		if !helpers.ValidarEmail(t.Email) {
			http.Error(w, "El email de usuario no es valido", 400)
			return
		}
	}

	status, err := database.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Error al modificar el registro del usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Error al modificar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
