package internal

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

func NewSenkyouServer(config Config, broker broker.Broker) SenkyouServer {
	return server{
		config: config,
		broker: broker,
	}
}

type server struct {
	config Config
	broker broker.Broker
}

func (s server) Start() {
	logger.Info("starting senkyou http server")
	fmt.Println(s.config.string())
	router := mux.NewRouter()
	router.HandleFunc("/", s.home)
	router.HandleFunc("/pub/{topic}/", s.pub)
	logger.Error("cannot start senkyou server", zap.Error(http.ListenAndServe(s.config.ListenAddr(), router)))
}

func (s server) pub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	topic := vars["topic"]
	logger.Debug("entering pub", zap.String("topic", topic))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.broker.Publish(topic, body)
	w.WriteHeader(http.StatusAccepted)
}

func (server) home(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("senkyou is up!\n"))
}
