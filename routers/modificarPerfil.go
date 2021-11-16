package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Luisul/curso-go-twitter/bd"
	"github.com/Luisul/curso-go-twitter/models"
)

/*ModificarPerfil modifica perfil usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar perfil"+err.Error(), 500)
		return
	}

	if !status {
		http.Error(w, "No se logró modificar la info del usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
