package internal

import (
	"errors"
	"github.com/abdelhamidbakhta/senkyou/internal/broker"
)

func NewBroker(config Config) (broker.Broker, error) {
	switch config.BrokerType {
	case broker.Nats:
		return broker.NewNatsBroker(config.NatsUrl), nil
	case broker.Kafka:
		return broker.NewKafkaBroker(config.KafkaUrl), nil
	default:
		return nil, errors.New("unsupported broker type")
	}
}
