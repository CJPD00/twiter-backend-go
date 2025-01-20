package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CJPD00/twiter-backend-go/database"
	"github.com/CJPD00/twiter-backend-go/jwt"
	"github.com/CJPD00/twiter-backend-go/models"
)

// Login handles user authentication by validating email and password from the request body.
// It checks if the user exists in the database and generates a JWT token for successful login.
// The token is returned in the response and set as a cookie for session management.
// If any error occurs during the process, an appropriate HTTP error is returned.
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en la peticion "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 || len(t.Password) == 0 {
		http.Error(w, "El email y la contraseña son requeridos", 400)
		return
	}

	documento, existe := database.IntentoLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario o contraseña incorrectos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "Error al generar el token "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
