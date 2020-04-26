package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/WhistleNewsBackend/src/app/model"
	"github.com/WhistleNewsBackend/src/app/ws"

	"github.com/bitly/go-nsq"
)

//ArticleViewsHandler distribute message
func (n *NSQ) ArticleViewsHandler(message *nsq.Message) error {
	message.Touch()
	fmt.Printf("Article viewd: %v", string(message.Body[:]))

	// Decode message
	var articleMes model.ArticleViewedMsg
	err := json.Unmarshal(message.Body, &articleMes)
	if err != nil {
		log.Println(err)
	}
	// Find Article
	article, err := n.Repo.GetArticle(articleMes.ID)
	if err != nil {
		log.Println(err)
	}
	res, err := json.Marshal(struct {
		Article *model.Article `json: "article"`
	}{
		Article: article,
	})
	if err != nil {
		log.Println(err)
	}

	n.Hub.Emit <- &ws.Emitter{Ids: []string{articleMes.Session}, Message: res}

	message.Finish()
	return nil
}
