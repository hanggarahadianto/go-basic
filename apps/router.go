package app

import (
	"origin/apps/controllers"

	"github.com/gorilla/mux"
)


func (server *Server) initalizeRoutes(){
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/",controllers.Home).Methods("GET")
}