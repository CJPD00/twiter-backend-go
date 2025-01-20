package routers

import (
	"encoding/json"
	"github.com/CJPD00/twiter-backend-go/database"
	"log"
	"net/http"
	"strconv"
)

func LeoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Debe enviar el parametro page como entero mayor a 0", http.StatusBadRequest)
		return
	}

	tweets, exito := database.LeoTweets(ID, page)
	log.Println(tweets)

	if !exito {
		http.Error(w, "Error al leer los tweets ", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}
