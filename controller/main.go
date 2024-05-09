package main

import (
	"log"
	server "v1"
)

func main(){
	server := new(server.Server)
	error := server.Run("8000")
	if error != nil {
		log.Fatal("error running http server")
	}
}