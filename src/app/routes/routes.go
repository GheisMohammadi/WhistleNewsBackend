package routes

import (
	"net/http"

	"github.com/WhistleNewsBackend/src/app/api"
	"github.com/WhistleNewsBackend/src/app/ws"
)

//Route instance
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//LoadRoutes initiates routes
func LoadRoutes(svc *api.API, websocket *ws.WebSocket) []Route {

	return []Route{
		Route{
			"WS",
			"GET",
			"/ws",
			websocket.ServeWs,
		},
		Route{
			"Index",
			"GET",
			"/",
			svc.Index,
		},
		Route{
			"AddArticle",
			"POST",
			"/article/add",
			svc.CreateArticle,
		},
		Route{
			"Statistics",
			"POST",
			"/statistics",
			svc.AddView,
		},
		Route{
			"GetStatistics",
			"GET",
			"/statistics/article_id/{id}",
			svc.GetArticle,
		},
	}

}
