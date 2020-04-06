package broker

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
