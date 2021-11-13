package middlew

import (
	"net/http"

	"github.com/Luisul/curso-go-twitter/routers"
)

/*ValidoJWT permiye validar el JWT en las peticiones*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(rw, "Error en el Token !!"+err.Error(), http.StatusBadRequest)
			return
		}
	}
}
