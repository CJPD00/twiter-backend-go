package models

import ()

// RespuestaLogin es la estructura de la respuesta del login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
