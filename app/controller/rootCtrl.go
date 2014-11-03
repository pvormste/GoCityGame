package controller

import (
	"net/http"
	"vormstein.eu/gocitygame/app"
	"vormstein.eu/gocitygame/app/views"
)

// Handler for "/"
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		views.GetTemplate("#index").ExecuteTemplate(w, "base", nil)
	default:
		views.GetTemplate("#error").ExecuteTemplate(w, "base", app.GetHTTPError(405))
	}
}

// Handler for "/register"
/*func HandleRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//
		default:
	}
}*/
