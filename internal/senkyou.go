package internal

import (
	"github.com/abdelhamidbakhta/senkyou/internal/broker"
	"github.com/abdelhamidbakhta/senkyou/internal/log"
	"go.uber.org/zap"
)

var logger = log.ForceGetLogger()

type Senkyou interface {
	Start()
}

func NewSenkyou(config Config, broker broker.Broker) (Senkyou, error) {
	return senkyou{
		config: config,
		broker: broker,
	}, nil
}

type senkyou struct {
	config Config
	broker broker.Broker
}

func (s senkyou) Start() {
	defer logger.Sync()
	err := s.broker.Subscribe(s.config.TopicIncomingRpcRequests, s.onIncomingRequest)
	if err != nil {
		logger.Error("failed to subscribe to incoming request topic", zap.Error(err))
	}
}

func (s senkyou) onIncomingRequest(request []byte) {
	requestBody := string(request)
	logger.Info("receiving new incoming request", zap.String("body", requestBody))
}
