package broker

const (
	Nats  = "nats"
	Kafka = "kafka"
)

type BrokerType string

type Broker interface {
	Publish(topic string, message []byte) error
	Subscribe(topic string, handler EventHandler) error
}

type EventHandler func([]byte)
