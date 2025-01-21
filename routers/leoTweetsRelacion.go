package routers

import (
	"encoding/json"
	"github.com/CJPD00/twiter-backend-go/database"
	"log"
	"net/http"
	"strconv"
)

func LeoTweetsRelacion(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	if len(page) < 1 {
		http.Error(w, "Debe enviar el parametro page", http.StatusBadRequest)
		return
	}

	pag, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Debe enviar el parametro page como entero mayor a 0", http.StatusBadRequest)
		return
	}

	if pag < 1 {
		http.Error(w, "Debe enviar el parametro page como entero mayor a 0", http.StatusBadRequest)
		return
	}

	result, status := database.LeoTweetsSeguidores(IDUsuario, pag)

	if !status {
		log.Println("erro en la operacion contra la base de datos")
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
