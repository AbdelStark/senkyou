package net

import (
	"fmt"
	"github.com/abdelhamidbakhta/senkyou/internal/broker"
	"github.com/abdelhamidbakhta/senkyou/internal/log"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

var (
	logger = log.ForceGetLogger()
)

type SenkyouServer interface {
	Start()
}

func NewSenkyouServer(listenAddr string, broker broker.Broker) SenkyouServer {
	return server{
		listenAddr: listenAddr,
		broker:     broker,
	}
}

type server struct {
	listenAddr string
	broker     broker.Broker
}

func (s server) Start() {
	logger.Info("starting senkyou http server")
	defer logger.Sync()
	router := mux.NewRouter()
	router.HandleFunc("/", s.home)
	router.HandleFunc("/pub/{topic}/", s.pub)
	router.HandleFunc("/sub/{topic}/", s.sub)
	logger.Error("cannot start senkyou server", zap.Error(http.ListenAndServe(s.listenAddr, router)))
}

func (s server) pub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	topic := vars["topic"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("failed to read request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.broker.Publish(topic, body)
	if err != nil {
		logger.Error("failed to publish message", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (s server) sub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	topic := vars["topic"]
	err := s.broker.Subscribe(topic, func(message []byte) {
		fmt.Println("new message")
		logger.Info("received message", zap.String("topic", topic), zap.String("message", string(message)))
	})
	if err != nil {
		logger.Error("failed to subscribe to topic", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (server) home(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("senkyou is up!\n"))
}
