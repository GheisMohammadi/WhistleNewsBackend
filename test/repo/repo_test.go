package repo_test

import (
	"testing"

	"WhistleNewsBackend/db"
	"WhistleNewsBackend/repo"

	config "WhistleNewsBackend/config"

	"gopkg.in/mgo.v2"
)

type Prep struct {
	Session    *mgo.Session
	Repo       *repo.Repo
	Collection string
}

func InitPrep(collection string) *Prep {
	mongo := &db.MONGO{
		Uri:      config.MongoURI,
		Database: config.DatabaseUserName,
	}
	mongo.Dial()
	dbM := &repo.Repo{
		Mongo: mongo,
	}
	return &Prep{
		Repo:       dbM,
		Collection: collection,
	}
}

func (prep *Prep) TearDown() {
	db, session := prepArticle.Repo.GetMgSession()
	defer session.Close()
	db.C(prep.Collection).RemoveAll(nil)
}

func TestMain(m *testing.M) {
	preps := []*Prep{
		prepArticle,
	}
	m.Run()
	for _, pre := range preps {
		pre.TearDown()
	}
}
