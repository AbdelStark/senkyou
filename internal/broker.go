package internal

import "errors"

const (
	Nats  = "nats"
	Kafka = "kafka"
)

type BrokerType string

type Broker interface {
	Publish(topic string, message []byte)
	Subscribe(topic string, handler EventHandler)
}

type EventHandler func([]byte) error

func NewBroker(config Config) (Broker, error) {
	switch config.BrokerType {
	case Nats:
		return newNatsBroker(config), nil
	case Kafka:
		return newKafkaBroker(config), nil
	default:
		return nil, errors.New("unsupported broker type")
	}
}
