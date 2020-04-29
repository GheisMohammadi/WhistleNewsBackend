package model_test

import (
	"WhistleNewsBackend/model"
	"testing"
	"time"
)

func TestCreatArticle(t *testing.T) {
	article := model.InitializeArticle()
	article.ID = "abcdefghi"
	err := article.AddView(time.Now(), 10)
	if err != nil {
		t.Error(err)
	}
}