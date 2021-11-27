package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Luisul/curso-go-twitter/bd"
)

func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar parametro page mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
	if !status {
		http.Error(w, "Error al leer usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
