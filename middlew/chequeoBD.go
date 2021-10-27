package middlew

import (
	"net/http"

	"github.com/Luisul/curso-go-twitter/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(rw, "Conexión perdida BD", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
