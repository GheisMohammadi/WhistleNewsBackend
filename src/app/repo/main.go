package repo

import (
	"gopkg.in/mgo.v2"

	"github.com/WhistleNewsBackend/src/app/db"
)

/**
* Repo hold different datasources
* Initialize Repo on each request
 */
type Repo struct {
	Mongo *db.MONGO
}

func (repo *Repo) GetMgSession() (*mgo.Database, *mgo.Session) {
	return repo.Mongo.GetSession()
}
