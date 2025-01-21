package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CJPD00/twiter-backend-go/database"
)

func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	typeUser := r.URL.Query().Get("type")
	search := r.URL.Query().Get("search")

	if len(typeUser) < 1 {
		http.Error(w, "Debe enviar el parametro type", http.StatusBadRequest)
		return
	}

	if len(page) < 1 {
		http.Error(w, "Debe enviar el parametro page", http.StatusBadRequest)
		return
	}

	pag, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Debe enviar el parametro page como entero mayor a 0", http.StatusBadRequest)
		return
	}

	result, status := database.LeoUsuariosTodos(IDUsuario, int64(pag), search, typeUser)

	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
