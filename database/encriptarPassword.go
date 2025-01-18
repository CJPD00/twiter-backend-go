package database

import (
	"golang.org/x/crypto/bcrypt"
)

func EncriptarPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(bytes), err
}
