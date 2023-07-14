package app

import (
	"net/http"
	"origin/apps/controllers"

	"github.com/gorilla/mux"
)


func (server *Server) initalizeRoutes(){
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/",controllers.Home).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}