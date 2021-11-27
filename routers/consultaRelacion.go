package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Luisul/curso-go-twitter/bd"
	"github.com/Luisul/curso-go-twitter/models"
)

/*ConsultaRelacion consulta la relación*/
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || !status {
		resp.Status = status
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
