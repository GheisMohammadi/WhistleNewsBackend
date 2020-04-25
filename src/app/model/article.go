package model

import (
	"net/http"
	"time"

	"github.com/mholt/binding"
	"gopkg.in/mgo.v2/bson"
)

//ArticleView object
type ArticleView struct {
	Reference string `bson:"reference" json:"reference" valid:"required"`
	Count     uint64 `bson:"count" json:"count" valid:"required"`
}

//Article object
type Article struct {
	ID        string         `bson:"_id" json:"_id" valid:"required"`
	Views     []*ArticleView `bson:"views" json:"views" valid:"required"`
	CreatedAt time.Time      `bson:"createdAt" json:"createdAt" valid:"required"`
	UpdatedAt time.Time      `bson:"updatedAt" json:"updatedAt" valid:"required"`
}

//InitializeArticle Initializes an article
func InitializeArticle() *Article {
	return &Article{
		ID:        bson.NewObjectId().Hex(),
		Views:     make([]*ArticleView, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

//AddView adds new new to an article
func (article *Article) AddView(ref string, count uint64) {

	newArticleView := &ArticleView{Reference: ref, Count: count}
	article.Views = append(article.Views, newArticleView)

}

//ArticleReq defines req for article
type ArticleReq struct {
	ID    string   `json:"id" xml:"id" form:"id" valid:"required"`
	Views []string `json:"views" xml:"views" form:"views" valid:"required"`
}

//FieldMap for mapping request
func (l *ArticleReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&l.ID:    "id",
		&l.Views: "views",
	}
}

//ArticleViewMes viewed message
type ArticleViewMes struct {
	ID      string `json:"id" xml:"id" form:"id" valid:"required"`
	Session string `json:"session" xml:"session" form:"session" valsession:"required"`
}
