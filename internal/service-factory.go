package internal

import (
	"errors"
	"github.com/abdelhamidbakhta/senkyou/internal/broker"
	"github.com/abdelhamidbakhta/senkyou/internal/config"
)

func NewBroker(cfg config.Config) (broker.Broker, error) {
	switch cfg.BrokerType {
	case config.BrokerNats:
		return broker.NewNatsBroker(cfg.NatsUrl)
	case config.BrokerKafka:
		return broker.NewKafkaBroker(cfg.KafkaUrl)
	default:
		return nil, errors.New("unsupported broker type")
	}
}
