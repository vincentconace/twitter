package bd

import (
	"context"
	"time"

	"github.com/vincentconace/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChequeoYaExisteUsuario esta funcion verifica si un usuario ya existe en la base de datos
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//se asigna a la variable db la base de datos "twittor"
	db := MongoCN.Database("twitter")
	//se asigna a al varaible col la coleccion de la base de datos
	col := db.Collection("usuarios")

	//cindicion para verificar en la coleccion si existe el email ingresado
	condicion := bson.M{"email": email}

	//variable de timo Usuario
	var resultado models.Usuario

	//se verifica la existenca del email
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
