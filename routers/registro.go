package routers

import (
	"encoding/json"
	"github.com/CJPD00/twiter-backend-go/database"
	"github.com/CJPD00/twiter-backend-go/models"
	"net/http"
)

// Registro permite crear un nuevo usuario en la base de datos.
// Recibe un JSON con los datos del usuario a registrar, verifica que
// el email no este duplicado, y que la contraseña tenga al menos 6
// caracteres. Si el registro es exitoso, se devuelve un status 201.
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en la peticion "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := database.ChequeoYaExisteUsuario(t.Email)

	if encontrado {
		http.Error(w, "El usuario ya existe", 400)
		return
	}

	_, status, err := database.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
