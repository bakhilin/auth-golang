package main

import (
	"log"
	server "v1"
	"v1/lib/handler"
)

func main(){
	handlers := new(handler.Handler)		
	server := new(server.Server)
	error := server.Run("8000", handlers.InitRoutes())
	if error != nil {
		log.Fatal("error running http server")
	}
}