package server

import (
	"github.com/WhistleNewsBackend/src/app/api"
	"github.com/WhistleNewsBackend/src/app/db"
	"github.com/WhistleNewsBackend/src/app/repo"
	"github.com/WhistleNewsBackend/src/app/routes"
	"github.com/WhistleNewsBackend/src/app/worker"
	"github.com/WhistleNewsBackend/src/app/ws"
	"github.com/WhistleNewsBackend/src/app/config"
	"github.com/gorilla/mux"
)

var (
	svc       *api.API
	websocket *ws.WebSocket
)

//NewRouter creates a router from all Routes
func NewRouter() *mux.Router {

	routes := routes.LoadRoutes(svc, websocket)
	router := mux.NewRouter().StrictSlash(true)
	router = router.PathPrefix(config.APIPrefix).Subrouter()
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

//Init initiates mongodb and NSQ
func Init() {
	//ws hub
	hub := ws.NewHub()
	go hub.Run()

	websocket = &ws.WebSocket{
		Hub: hub,
	}

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
	err := nSQ.CreateHandler(nSQ.ArticleViewsHandler, "articleviews", "statistics")
	if err != nil {
		panic(err)
	}

	// API share the same connection pool
	svc = &api.API{
		Repo: sharedRepo,
		Nsq:  nSQ,
	}

}
