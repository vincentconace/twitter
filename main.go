package main

import (
	"log"

	"github.com/vincentconace/twitter/bd"
	"github.com/vincentconace/twitter/handlers"
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Handlers()
}
