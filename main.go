package main

import (
	"fmt"
	"log"
	"net/http"
	app "origin/apps"

	"github.com/gorilla/mux"
)

var pl = fmt.Println


func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func handleRequest(){

	router := mux.NewRouter().StrictSlash(true)



	router.HandleFunc("/", helloWorld).Methods("GET")

	// router.HandleFunc("/users", allUsers).Methods("GET")
	// router.HandleFunc("/user", createUsers).Methods("POST")
	// router.HandleFunc("/user/{id}", updateUser).Methods("PATCH")
	// router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	

	log.Fatal(http.ListenAndServe(":8080", router))

	


}


func main(){

	x := 10
	if x > 5 {
		pl("x lebih besar")
	}else{
		pl("tolol")
	}

	book:= "the color of magic"
	pl(book)

	greetUser()

	// pl("go run in 8080")
	// handleRequest()

	app.Run()

	
}

func greetUser() {
	pl("Welcome to our Lontong")
}