package main

import (
	"log"

	"github.com/vincentconace/twittor/bd"
	"github.com/vincentconace/twittor/handlers"
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Handlers()
}
