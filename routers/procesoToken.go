package routers

import (
	"errors"
	"strings"

	"github.com/CJPD00/twiter-backend-go/database"
	"github.com/CJPD00/twiter-backend-go/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email usado en todos los endpoints
var Email string

// ID usado en todos los endpoints
var IDUsuario string

// ProcesoToken es un middleware que permite verificar el token de un usuario
// para asi poder acceder a los endpoints, si el token es invalido, se
// devuelve un error 400.
// Recibe el token en formato string y devuelve un modelo de Claim, un booleano
// indicando si el token es valido o no, el ID del usuario en formato string y
// un error.
func ProcesoToken(t string) (*models.Claim, bool, string, error) {

	miClave := []byte("c00120262364")
	claims := &models.Claim{}

	splitToken := strings.Split(t, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}

	t = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := database.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil

	}

	if !tkn.Valid {
		return claims, false, "", errors.New("token invalido")
	}

	return claims, false, "", err

}
