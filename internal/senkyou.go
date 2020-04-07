package internal

import (
	"github.com/abdelhamidbakhta/senkyou/internal/broker"
	"github.com/abdelhamidbakhta/senkyou/internal/log"
	"github.com/abdelhamidbakhta/senkyou/internal/net"
	"go.uber.org/zap"
)

var logger *zap.Logger

type Senkyou interface {
	Start()
}

func NewSenkyou(config Config, broker broker.Broker) (Senkyou, error) {
	logger = log.GetLoggerWithLevel(config.LogLevel.ZapLevel)
	return senkyou{
		config:    config,
		broker:    broker,
		rpcClient: net.NewRpcClient(config.RpcUrl),
	}, nil
}

type senkyou struct {
	config    Config
	broker    broker.Broker
	rpcClient net.RpcClient
}

func (s senkyou) Start() {
	defer logger.Sync()
	err := s.broker.Subscribe(s.config.TopicIncomingRpcRequests, s.onIncomingRequest)
	if err != nil {
		logger.Error("failed to subscribe to incoming request topic", zap.Error(err))
	}
}

func (s senkyou) onIncomingRequest(request []byte) {
	response, err := s.rpcClient.Call(request)
	if err != nil {
		s.handleError(err)
		return
	}
	s.handleError(s.broker.Publish(s.config.TopicOutgoingRpcResponses, response))
}

func (s senkyou) handleError(err error) {
	if err != nil {
		logger.Error("error occurred", zap.Error(err))
		_ = s.broker.Publish(s.config.TopicErrors, []byte(err.Error()))
	}
}
