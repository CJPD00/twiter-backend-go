package jwt

import (
	"os"
	"time"

	"github.com/CJPD00/twiter-backend-go/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// GeneroJWT generates a JWT token for the given user, including the user's email, name, last name, birthdate, biography, website, and ID.
// The token is signed with the secret key "c00120262364" and is valid for 24 hours.
func GeneroJWT(t models.Usuario) (string, error) {
	godotenv.Load()
	clave := os.Getenv("SECRET_KEY")
	miClave := []byte(clave)
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"sitio_web":        t.SitioWeb,
		"id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
