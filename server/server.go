package server

import (
	"WhistleNewsBackend/api"
	"WhistleNewsBackend/db"
	"WhistleNewsBackend/repo"
	"WhistleNewsBackend/routes"
	"WhistleNewsBackend/worker"
	"WhistleNewsBackend/ws"
	"WhistleNewsBackend/config"
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
	err := nSQ.CreateHandler(nSQ.ArticleViewsHandler, "article_viewd", "statistics")
	if err != nil {
		println("initialization of NSQ failed! check if NSQ services is running on your local system")
		panic(err)
	}

	// API share the same connection pool
	svc = &api.API{
		Repo: sharedRepo,
		Nsq:  nSQ,
	}

}
