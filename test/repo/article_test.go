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
	result, err := prepArticle.Repo.GetJob(article.ID)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
