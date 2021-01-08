package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//==============================
//Conctando a la base de datos
//==============================

//MongoCN es el objeto que conecta a la base de datos
var MongoCN = ConnectBD()

//URI para conecta a la base de datos
var clientOptions = options.Client().ApplyURI("mongodb+srv://vincent:v-21032991@twittor.8yvhb.mongodb.net/twittor?retryWrites=true&w=majority")

//ConnectBD esta funcion nos conecta con la base de datos
func ConnectBD() *mongo.Client {
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
	log.Println("Conections DB")
	return client
}

func ChequeoConection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
