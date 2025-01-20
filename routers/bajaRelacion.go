package routers

import (
	"github.com/CJPD00/twiter-backend-go/database"
	models "github.com/CJPD00/twiter-backend-go/models"
	"net/http"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("id")) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = r.URL.Query().Get("id")

	status, err := database.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Error al borrar relacion "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado borrar la relacion", 400)
		return
	}

	w.WriteHeader(http.StatusOK)
}
