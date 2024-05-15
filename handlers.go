package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r * http.Request){
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	if validateUser(r) {
		fmt.Println("SUCCESSFUL LOGIN!")	
	} else {
		fmt.Println("Sorry, permision denied!")
	}
}

func validateUser(r * http.Request) bool {
	if r.Form["login"][0] == Admin.getLogin() &&
			r.Form["password"][0] == Admin.getPassword() {
		return true
	}
	return false
}

func signUp(w http.ResponseWriter, r * http.Request) {
	// need to write it
} 