package workers

import (
	"WhistleNewsBackend/repo"
	"WhistleNewsBackend/ws"

	"github.com/bitly/go-nsq"
)

/*
NSQ construction
*/
type NSQ struct {
	Uri    string
	Config *nsq.Config
	Hub    *ws.Hub
	Repo   *repo.Repo
}

func InitNSQ(uri string, hub *ws.Hub, repo *repo.Repo) *NSQ {
	config := nsq.NewConfig()
	return &NSQ{uri, config, hub, repo}
}

//Establish producer
func (n *NSQ) CreateProducer() (*nsq.Producer, error) {
	w, err := nsq.NewProducer(n.Uri, n.Config)
	return w, err
}

//Establish consumer
func (n *NSQ) CreateConsumer(topic string, ch string) (*nsq.Consumer, error) {
	w, err := nsq.NewConsumer(topic, ch, n.Config)
	return w, err
}

func (n *NSQ) CreateHandler(handler func(*nsq.Message) error, topic string, ch string) error {
	consumer, err := n.CreateConsumer(topic, ch)
	consumer.AddHandler(nsq.HandlerFunc(handler))
	err = consumer.ConnectToNSQD(n.Uri)
	if err != nil {
		return err
	}
	return nil
}
