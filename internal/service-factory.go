package internal

import (
	"errors"
	"github.com/abdelhamidbakhta/senkyou/internal/broker"
	"github.com/abdelhamidbakhta/senkyou/internal/config"
)

func NewBroker(cfg config.Config) (broker.Broker, error) {
	var coreBroker broker.Broker
	var err error
	switch cfg.BrokerType {
	case config.BrokerNats:
		coreBroker, err = broker.NewNatsBroker(cfg.NatsUrl)
	case config.BrokerKafka:
		coreBroker, err = broker.NewKafkaBroker(cfg.KafkaUrl)
	default:
		return nil, errors.New("unsupported broker type")
	}
	if cfg.ApmEnabled {
		coreBroker = broker.NewApmBroker(coreBroker)
	}
	return coreBroker, err
}
