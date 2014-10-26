package main

import (
	"log"
	"vormstein.eu/gocitygame/util"
)

func main() {
	// Establish Conn
	session, _ := util.DB.GetCollection("")
	defer session.Close()

	// Drop the Database
	session.DB(util.Conf.DB_Name).DropDatabase()

	// Feedback
	log.Println("Dropped Database:", util.Conf.DB_Name)
}
