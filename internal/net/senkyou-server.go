package net

import (
	"github.com/abdelhamidbakhta/senkyou/internal/broker"
	"github.com/abdelhamidbakhta/senkyou/internal/config"
	"github.com/abdelhamidbakhta/senkyou/internal/log"
	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"net/http"
)

var (
	logger *zap.Logger
)

type SenkyouServer interface {
	Start()
}

func NewSenkyouServer(config config.Config, broker broker.Broker, logLevel zapcore.Level) SenkyouServer {
	logger = log.GetLogger(config)
	return server{
		config: config,
		broker: broker,
	}
}

type server struct {
	config config.Config
	broker broker.Broker
}

func (s server) Start() {
	logger.Info("starting senkyou http server")
	defer logger.Sync()
	router := mux.NewRouter()
	if s.config.ApmEnabled {
		router.Use(apmgorilla.Middleware())
	}
	router.HandleFunc("/", s.home)
	router.HandleFunc("/pub/{topic}/", s.pub)
	router.HandleFunc("/sub/{topic}/", s.sub)
	logger.Error("cannot start senkyou server", zap.Error(http.ListenAndServe(s.config.ListenAddr(), router)))
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
	err = s.broker.Publish(r.Context(), topic, body)
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
	traceContextFields := apmzap.TraceContext(r.Context())
	logger.With(traceContextFields...).Debug("handling home request")
	_, _ = w.Write([]byte("senkyou is up!\n"))
}
