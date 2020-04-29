package db_test

import (
	"testing"

	"WhistleNewsBackend/api"
	"WhistleNewsBackend/db"
	"WhistleNewsBackend/ws"

	config "WhistleNewsBackend/config"

	"github.com/gorilla/mux"
)

var (
	svc *api.API
	r   *mux.Router
)

func TestMongo(t *testing.T) {
	//ws hub
	hub := ws.NewHub()
	go hub.Run()

	// Initialize connetion pool with mongo
	mongo := &db.MONGO{
		Uri:      config.MongoURI,
		Database: config.DataBaseName,
	}
	mongo.Dial()

	db, session := mongo.GetSession()
	db.C("articles").RemoveAll(nil)
	session.Close()
}
