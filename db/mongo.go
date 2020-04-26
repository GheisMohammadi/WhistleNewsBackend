package db

import (
	"errors"

	"gopkg.in/mgo.v2"
)

/*
Mongo construction
*/
//Mongo struct hold uri based on environment
type MONGO struct {
	Uri      string
	Database string
	Session  *mgo.Session
}

//Establish connection to mongodb
func (mongo *MONGO) Dial() {
	session, err := mgo.Dial(mongo.Uri)
	if err != nil {
		panic(err)
	}
	mongo.Session = session
}

// Get session
func (mongo *MONGO) GetSession() (*mgo.Database, *mgo.Session) {
	if mongo.Session == nil {
		panic(errors.New("Db session not exist"))
		return nil, nil
	}

	newSession := mongo.Session.Copy()
	db := newSession.DB(mongo.Database)

	return db, newSession
}
