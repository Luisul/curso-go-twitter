package bd

import (
	"context"
	"time"

	"github.com/Luisul/curso-go-twitter/models"
)

/*BorroRelacion crea relacion tweets*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, err
}
