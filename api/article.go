package api

import (
	"encoding/json"
	"net/http"
	"time"

	"WhistleNewsBackend/model"
	"WhistleNewsBackend/utils"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

type Attributes struct {
	Views []*model.ArticleView `bson:"count" json:"count" valid:"required"`
}

type Data struct {
	ID             string      `bson:"article_id" json:"article_id" valid:"required"`
	Type           string      `bson:"type" json:"type" valid:"required"`
	DataAttributes *Attributes `bson:"attributes" json:"attributes" valid:"required"`
}
type StaticsResponse struct {
	Result *Data `bson:"data" json:"data" valid:"required"`
}

// GetArticle fetchs article
/*
* GET /counter/v1/statistics/article_id/:id
 */
func (api *API) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID := vars["id"]
	article, err := api.Repo.GetArticle(articleID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-type", "application/json")

	layout := "2006-01-02T15:04" //:05.000Z"
	viewsMap := make(map[string]uint64)

	for _, articleView := range article.Views {
		t,_:=time.ParseInLocation(layout, articleView.Reference, time.Local )
		sinceStr := utils.TimeSince(t)
		viewsMap[sinceStr]+=articleView.Count
	}
	
	var views []*model.ArticleView
	
	for key, val := range viewsMap {
		v := &model.ArticleView{
			Reference: key,
			Count: val,
		}
		views=append(views,v)
	}
	
	result := &Data{
		ID:   article.ID,
		Type: "statistics_article_view_count",
		DataAttributes: &Attributes{
			Views: views,
		},
	}
	resp := StaticsResponse{
		Result: result,
	}

	json.NewEncoder(w).Encode(resp)
}

//CreateArticle creates new article
/**
* POST /counter/v1/article/add
 */
func (api *API) CreateArticle(w http.ResponseWriter, r *http.Request) {
	/*
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		token := cookie.Value
	*/
	token := "abcdefghijklmnopqrstuvwxyz"

	articleReq := new(model.ArticleReq)
	errs := binding.Bind(r, articleReq)
	if errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errs.Error()))
		return
	}

	if res, err := govalidator.ValidateStruct(articleReq); res == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	article := model.InitializeArticle()
	article.ID = articleReq.ID

	// Persisting skillset
	errArticleCreation := api.Repo.CreateArticle(article)
	if errArticleCreation != nil {
		w.WriteHeader(500)
		return
	}
	// Publish message to nsq
	// Token is assigned from node frontend
	mes := &model.ArticleCreatedMsg{
		ID:      article.ID,
		Session: token,
	}
	mesStr, _ := json.Marshal(mes)
	go api.PublishNSQMes("article_created", []byte(mesStr))

	// Response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

//AddView adds new view to article statistics
/**
* POST /counter/v1/statistics
 */
func (api *API) AddView(w http.ResponseWriter, r *http.Request) {
	/*
		cookie, err := r.Cookie("token")
		println("cookie:",err)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		token := cookie.Value
	*/
	token := "abcdefghijklmnopqrstuvwxyz"

	articleViewReq := new(model.ArticleViewReq)
	errs := binding.Bind(r, articleViewReq)
	if errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errs.Error()))
		return
	}

	if res, err := govalidator.ValidateStruct(articleViewReq); res == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	articleID := articleViewReq.ID

	// Persisting skillset
	errAddView := api.Repo.AddViewToArticle(articleID)
	if errAddView != nil {
		w.WriteHeader(500)
		return
	}

	// Publish message to nsq
	// Token is assigned from node frontend
	mes := &model.ArticleViewedMsg{
		ID:      articleID,
		Session: token,
	}
	mesStr, _ := json.Marshal(mes)
	go api.PublishNSQMes("article_viewed", []byte(mesStr))

	// Response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		ID string `json: "id"`
	}{
		ID: articleID,
	})
}
