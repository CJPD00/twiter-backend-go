package database

import (
	"github.com/CJPD00/twiter-backend-go/models"
	"golang.org/x/crypto/bcrypt"
)

// IntentoLogin permite verificar si un usuario con el email y password dados
// existe en la base de datos y si coincide. Retorna el modelo de usuario
// encontrado y un booleano indicando si el usuario y su contrasen ̃a coinciden
// o no. Si el usuario no existe, se devuelve un usuario vacio y false.
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)

	if !encontrado {
		return usuario, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usuario, false
	}

	return usuario, true

}
