package jwt

import (
	"time"

	"github.com/Luisul/curso-go-twitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT genera el encriptado con JWT*/
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MastersdelDesarrollo")

	payload := jwt.MapClaims{
		"email":              t.Email,
		"nombre":             t.Nombre,
		"apellidos":          t.Apellidos,
		"fecha_nacimimiento": t.FechaNacimiento,
		"biografia":          t.Biografia,
		"ubicacion":          t.SitioWeb,
		"_id":                t.ID.Hex(),
		"exp":                time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
