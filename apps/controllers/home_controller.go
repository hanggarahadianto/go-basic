package controllers

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"
)

var pf = fmt.Fprintf

func Home(w http.ResponseWriter, r *http.Request){

	render:= render.New(render.Options{
		Layout: "layout",
	})

	_ = render.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title": "Home Title",
		"body": "Home Description",
	})

	pf(w, "Welcome to out home TOKO")
}