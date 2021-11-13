package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Luisul/curso-go-twitter/bd"
)

/*VerPerfil consulta la información del perfil usuario*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {

		http.Error(w, "Ocurrió un error al intentar consultar el registro"+err.Error(), 400)
		return
	}

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
