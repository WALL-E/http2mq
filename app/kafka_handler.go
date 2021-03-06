package app

import (
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Kafka struct {
	Topic string
}

func (k *Kafka) DoGet(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("{\"messge\":\"ok\"}"))
}

func (k *Kafka) DoPost(res http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s: res:%s", req.RequestURI, err.Error())
		res.Write([]byte(""))
		return
	}

	Conf.SyncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: k.Topic,
		Value: sarama.ByteEncoder(b),
	})

	//
	//Conf.AsyncProducer.Input() <- &sarama.ProducerMessage{
	//	Topic: "http2mq",
	//	Value: sarama.ByteEncoder(b),
	//}

	res.Write([]byte(""))
}

func (k *Kafka) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k.Topic = vars["topic"]

	if !CheckTopic(k.Topic) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"messge\":\"无效的topic\"}"))
	}

	if !CheckAuth(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case http.MethodPost, http.MethodPut:
		k.DoPost(w, r)
	case http.MethodGet:
		k.DoGet(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(""))
	}
}

func NewKafka() http.Handler {
	return &Kafka{}
}
