package main

import (
	"fmt"
	"net/http"
)

var pf = fmt.Fprintf

func allUsers(w http.ResponseWriter , r* http.Request){
	pf(w, "allUsers")
}

func createUsers(w http.ResponseWriter, r *http.Request){
	pf(w, "userCreated")
}

func updateUser(w http.ResponseWriter, r *http.Request){
	pf(w, "user updated")
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	pf(w, "user deleted")
}