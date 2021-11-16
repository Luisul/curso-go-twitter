package routers

import (
	"net/http"

	"github.com/Luisul/curso-go-twitter/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar parametro id", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)

	if err != nil {
		http.Error(w, "Error al eliminar el tweet"+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
