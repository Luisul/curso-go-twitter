package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Luisul/curso-go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil busca perfil en la BD*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""

	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}

	return perfil, nil
}
