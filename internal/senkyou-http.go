package internal

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

var (
	logger, _ = zap.NewDevelopment()
)

type SenkyouServer interface {
	Start()
}

func NewSenkyouServer(config Config) SenkyouServer {
	return server{
		config: config,
	}
}

type server struct {
	config Config
}

func (s server) Start() {
	logger.Info("starting senkyou http server")
	fmt.Println(s.config.string())
	router := mux.NewRouter()
	router.HandleFunc("/", s.home)
	logger.Error("cannot start senkyou server", zap.Error(http.ListenAndServe(s.config.ListenAddr(), router)))
}

func (server) home(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("senkyou is up!\n"))
}
