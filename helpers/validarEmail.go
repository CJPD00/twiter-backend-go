package helpers

import (
	"regexp"
)

// ValidarEmail valida si un email es correcto
// Parte local (antes del @)

// Debe contener solo letras minúsculas (a-z)
// Puede contener dígitos (0-9)
// Puede contener puntos (.)
// Puede contener guiones bajos (_)
// Puede contener porcentajes (%)
// Puede contener signos más (+)
// Puede contener guiones (-)
// Debe tener al menos un carácter
// Parte del dominio (después del @)

// Debe contener solo letras minúsculas (a-z)
// Puede contener dígitos (0-9)
// Puede contener puntos (.)
// Puede contener guiones (-)
// Debe tener al menos un carácter
// Dominio de nivel superior

// Debe contener solo letras minúsculas (a-z)
// Debe tener entre 2 y 4 caracteres
// Otros

// La dirección de correo electrónico debe tener un formato válido, con una parte local, un dominio y un dominio de nivel superior.
// No se permiten espacios en blanco ni otros caracteres especiales.
// Ejemplos de correos correctos:
// ejemplo@ejemplo.com
// ejemplo@sub.ejemplo.com
// ejemplo.ejemplo@sub.ejemplo.com
// ejemplo-ejemplo@sub.ejemplo.com
// ejemplo_ejemplo@sub.ejemplo.com
func ValidarEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}
