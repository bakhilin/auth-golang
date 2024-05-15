package main 

import (
	"net/http"
)

const(
	host = "127.0.0.1"
)

var Admin = User {
	"ojito",
	"35791399",
	true,
}

func main(){
	server := http.Server{Addr: host+":8000"}
	http.HandleFunc("/login", login)
	http.HandleFunc("sign-up", signUp)
	server.ListenAndServe()

}