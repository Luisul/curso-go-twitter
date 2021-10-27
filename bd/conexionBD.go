package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN objeto de conexion*/
var MongoCN = ConectarBD()

/*clientOptions*/
var clientOptions = options.Client().ApplyURI("mongodb+srv://luribe:Luribe2022@cluster0.elm1s.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConectarBD Es la funcion para conectar la BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa BD")
	return client
}

/*ChequeoConnection ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
