package gobase

import (
	"github.com/nelsongp/gobase/bd"
	"github.com/nelsongp/gobase/handlers"
	"log"
)

func main(){
	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexion a la bd")
		return
	}
	handlers.Manejadores()
}