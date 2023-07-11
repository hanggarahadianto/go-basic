package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var pl = fmt.Println


func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func handleRequest(){

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))



}


func main(){
	pl("go run in 8080")
	handleRequest()
}