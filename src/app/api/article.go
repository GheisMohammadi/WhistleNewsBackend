package api

import (
	"github.com/WhistleNewsBackend/src/app/model"
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

// GetArticle fetchs article
/*
* GET /api/v1/statistics/article_id/:id
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
	json.NewEncoder(w).Encode(article)
}

//CreateArticle creates new article
/**
* POST /api/v1/article/add
*/
func (api *API) CreateArticle(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	token := cookie.Value

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
	//views := make([]*model.ArticleView, 0)

	// Persisting skillset
	err = api.Repo.CreateArticle(article)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// Publish message to nsq
	// Token is assigned from node frontend
	mes := &model.ArticleViewMes{
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