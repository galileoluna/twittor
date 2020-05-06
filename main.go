package main

import (
	"log"
	"quickstart/bd"
	"quickstart/handlers"

	"github.com/galileoluna/twittor/bd"
	"github.com/galileoluna/twittor/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()

}
