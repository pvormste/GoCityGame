package controller

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "%s", "Hello World at Root!")
	default:
		fmt.Fprintf(w, "%s", "Request not allowed")
	}
}
