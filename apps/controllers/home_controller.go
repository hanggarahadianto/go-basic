package controllers

import (
	"fmt"
	"net/http"
)

var pf = fmt.Fprintf

func Home(w http.ResponseWriter, r *http.Request){
	pf(w, "Welcome to out home TOKO")
}