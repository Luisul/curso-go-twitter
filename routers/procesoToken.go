package routers

import (
	"errors"
	"strings"

	"github.com/Luisul/curso-go-twitter/bd"
	"github.com/Luisul/curso-go-twitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email usuario*/
var Email string

/*IDUsuario usuario para usar en los EndPoints*/
var IDUsuario string

/*ProcesoToken proceso token extraer sus valores*/
func ProcesoToken(tk string) (*models.Claims, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo")
	claims := &models.Claims{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}
	return claims, false, string(""), err
}
