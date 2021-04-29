package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://twittor:P3r5s1a8191@cluster0-ruzay.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConectarBD agregar siempre conementario*/
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

	log.Println("Conexion Exitosa con la BD")
	return client
}

//ChequeoConnection es ping a la base
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}