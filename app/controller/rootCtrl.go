package controller

import (
	"fmt"
	"net/http"
	"vormstein.eu/gocitygame/app/views"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		views.Templates["index.html"].ExecuteTemplate(w, "base", nil)
	default:
		fmt.Fprintf(w, "%s", "Request not allowed")
	}
}
