package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"vormstein.eu/gocitygame/app/controller"
	"vormstein.eu/gocitygame/util"
)

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", controller.HandleRoot)

	// Static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Print an information
	log.Println("Listening at Port", util.Conf.Port)

	// Serve!
	http.Handle("/", router)
	http.ListenAndServe(util.Conf.Port, nil)
}
