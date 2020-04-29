package repo_test

import (
	"testing"

	"WhistleNewsBackend/model"
)

var (
	prepArticle *Prep
)

func init() {
	prepArticle = InitPrep("articles")
}

func TestCreatArticle(t *testing.T) {
	article := model.InitializeArticle()
	err := prepArticle.Repo.CreateArticle(article)
	if err != nil {
		t.Error(err)
	}
}

func TestGetArticle(t *testing.T) {
	article := model.InitializeArticle()
	db, session := prepArticle.Repo.GetMgSession()
	defer session.Close()
	db.C("articles").Insert(article)
	result, err := prepArticle.Repo.GetArticle(article.ID)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestAddViewToArticle(t *testing.T) {
	article := model.InitializeArticle()
	db, session := prepArticle.Repo.GetMgSession()
	defer session.Close()
	db.C("articles").Insert(article)
	errAddView := prepArticle.Repo.AddViewToArticle(article.ID)
	if errAddView != nil {
		t.Error(errAddView)
	}
	result, err := prepArticle.Repo.GetArticle(article.ID)
	if err != nil {
		t.Error(err)
	}
	if (len(result.Views)<=0){
		t.Error(err)
	}
	t.Log(result)
}
