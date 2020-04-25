package repo

import (
	"github.com/WhistleNewsBackend/src/app/model"
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
	err := db.C("articles").FindId(id).One(&article)
	if err != nil {
		return &model.Article{}, err
	}
	return &article, nil
}