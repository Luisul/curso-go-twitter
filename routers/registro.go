package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Luisul/curso-go-twitter/bd"
	"github.com/Luisul/curso-go-twitter/models"
)

/*Registro funcion para crear en la BD los usuarios*/
func Registro(rw http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(rw, "Email del usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(rw, "Especificar contraseña al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(rw, "Ya existe un usuario con el email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(rw, "No se ha logrado registrar el usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(rw, "Ocurrió un error al intentar registrar el usuario", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)

}
