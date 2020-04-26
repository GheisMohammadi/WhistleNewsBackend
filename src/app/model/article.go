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
func (article *Article) AddView(ref time.Time, count uint64) error{

	t:= ref.Format("2006-01-02T15:04")
	n:=len(article.Views)
	if (n>0 && article.Views[n-1].Reference == t){
		article.Views[n-1].Count += count
	}else{
		newArticleView := &ArticleView{Reference: t, Count: count}
		article.Views = append(article.Views, newArticleView)
	}
	return nil
}

//ArticleReq defines req for article creation
type ArticleReq struct {
	ID string `json:"id" xml:"id" form:"id" valid:"required"`
}

//FieldMap for mapping request to article req
func (ar *ArticleReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&ar.ID: "id",
	}
}

//ArticleViewReq defines request for new view for specific article
type ArticleViewReq struct {
	ID string `json:"id" xml:"id" form:"id" valid:"required"`
}

//FieldMap for mapping request to view req
func (ar *ArticleViewReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&ar.ID: "id",
	}
}

//ArticleCreatedMsg created message
type ArticleCreatedMsg struct {
	ID    string `json:"id" xml:"id" form:"id" valid:"required"`
	Session string `json:"session" xml:"session" form:"session" valsession:"required"`
}

//ArticleViewedMsg viewed message
type ArticleViewedMsg struct {
	ID    string `json:"id" xml:"id" form:"id" valid:"required"`
	Session string `json:"session" xml:"session" form:"session" valsession:"required"`
}
