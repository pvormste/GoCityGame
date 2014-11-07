package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"vormstein.eu/gocitygame/app"
	"vormstein.eu/gocitygame/app/controller"
)

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", controller.HandleRoot)
	router.HandleFunc("/register", controller.HandleRegister)

	// Static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Print an information
	log.Println("Listening at Port", app.Config.Port)

	// Serve!
	http.Handle("/", router)
	http.ListenAndServe(app.Config.Port, nil)
}
