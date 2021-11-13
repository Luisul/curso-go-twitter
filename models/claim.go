package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claims estructura para manejar el JWT*/
type Claims struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
