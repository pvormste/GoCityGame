package app

import (
	"gopkg.in/mgo.v2"
)

// The App Configuration struct
type AppConfiguration struct {
	Port    string
	DB_URL  string
	DB_Name string
	Tmpl    string
	DB      *DatabaseConfiguration
}

// The Database Configuration struct
type DatabaseConfiguration struct {
	session *mgo.Session
}

// ### Private

// Get Session or create one
func (db *DatabaseConfiguration) getSession() *mgo.Session {
	if db.session == nil {
		var err error
		db.session, err = mgo.Dial(Config.DB_URL)
		if err != nil {
			panic(err)
		}
	}

	return db.session.Clone()
}

// ### Public

// Return the desired collection
func (db *DatabaseConfiguration) GetCollection(coll string) (*mgo.Session, *mgo.Collection) {
	s := db.getSession()

	if coll == "" {
		return s, nil
	}

	return s, s.DB(Config.DB_Name).C(coll)
}
