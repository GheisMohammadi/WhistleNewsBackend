package repo

import (
	"time"

	"WhistleNewsBackend/model"
	"gopkg.in/mgo.v2/bson"
)

//CreateArticle creates new article in db
func (repo *Repo) CreateArticle(article *model.Article) error {
	db, session := repo.GetMgSession()
	defer session.Close()
	err := db.C("articles").Insert(article)
	return err
}

//GetArticle fetchs article from db
func (repo *Repo) GetArticle(id string) (*model.Article, error) {
	db, session := repo.GetMgSession()
	defer session.Close()
	var article model.Article
	err := db.C("articles").Find(bson.M{"_id": id}).One(&article)
	if err != nil {
		return &model.Article{}, err
	}
	return &article, nil
}

//AddViewToArticle adds new view to article in db
func (repo *Repo) AddViewToArticle(id string) error {
	db, session := repo.GetMgSession()
	defer session.Close()
	var article model.Article
	err := db.C("articles").Find(bson.M{"_id": id}).One(&article)
	if err != nil {
		return err
	}
	article.AddView(time.Now(), 1)
	article.UpdatedAt = time.Now()
	errUpdate := db.C("articles").Update(bson.M{"_id": id}, article)

	return errUpdate
}