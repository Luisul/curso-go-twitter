package routers

import (
	"net/http"

	"github.com/Luisul/curso-go-twitter/bd"
	"github.com/Luisul/curso-go-twitter/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio eliminar", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Occurrio un error al eliminar la relación"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se logró eliminar la relación", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
