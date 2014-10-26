package util

import (
	"gopkg.in/mgo.v2"
)

// Database struct
type Database struct {
	session *mgo.Session
}

// ## Private

// Get Session or create one
func (db *Database) getSession() *mgo.Session {
	if db.session == nil {
		var err error
		db.session, err = mgo.Dial(Conf.DB_URL)
		if err != nil {
			panic(err)
		}
	}

	return db.session.Clone()
}

// ## Public

// Return the desired collection
func (db *Database) GetCollection(coll string) (*mgo.Session, *mgo.Collection) {
	s := db.getSession()

	if coll == "" {
		return s, nil
	}

	return s, s.DB(Conf.DB_Name).C(coll)
}

// The instance
var DB Database

func init() {
	DB = Database{}
}
