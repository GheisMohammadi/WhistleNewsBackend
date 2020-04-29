package api_test

import (
	"net/http"
	"testing"

	"WhistleNewsBackend/api"
	"WhistleNewsBackend/db"
	"WhistleNewsBackend/model"
	"WhistleNewsBackend/repo"
	"WhistleNewsBackend/worker"
	"WhistleNewsBackend/ws"

	config "WhistleNewsBackend/config"

	"github.com/gorilla/mux"
)

var (
	svc       *api.API
	articleId string
	r         *mux.Router
)

func init() {
	//ws hub
	hub := ws.NewHub()
	go hub.Run()

	// Initialize connetion pool with mongo
	mongo := &db.MONGO{
		Uri:      config.MongoURI,
		Database: config.DataBaseName,
	}
	mongo.Dial()
	sharedRepo := &repo.Repo{
		Mongo: mongo,
	}

	//NSQ
	nSQ := workers.InitNSQ(config.NSQURL, hub, sharedRepo)
	err := nSQ.CreateHandler(nSQ.ArticleViewsHandler, "article_viewd", "statistics")
	if err != nil {
		panic(err)
	}

	// API share the same connection pool
	svc = &api.API{
		Repo: sharedRepo,
		Nsq:  nSQ,
	}

	// Create route with context
	r = mux.NewRouter()
	r.HandleFunc("/statistics/article_id/{id}", svc.GetArticle)

	http.Handle("/", r)

}

func SetUp() {
	article := model.InitializeArticle()
	articleId = article.ID

	db, session := svc.Repo.GetMgSession()
	defer session.Close()
	db.C("articles").Insert(&article)
}

func TearDown(collection string) {
	db, session := svc.Repo.GetMgSession()
	db.C(collection).RemoveAll(nil)
	session.Close()
}

func TestMain(m *testing.M) {
	SetUp()
	TearDown("articles")
}
