package api

import (
	"log"

	"WhistleNewsBackend/repo"
	"WhistleNewsBackend/worker"
)

type API struct {
	Repo *repo.Repo
	Nsq  *workers.NSQ
}

func (api *API) PublishNSQMes(topic string, mes []byte) {
	producer, err := api.Nsq.CreateProducer()
	if err != nil {
		log.Println(err)
	}
	err = producer.Publish(topic, mes)
	if err != nil {
		log.Println(err)
	}
	producer.Stop()
}
