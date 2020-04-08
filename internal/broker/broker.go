package broker



type Broker interface {
	Publish(topic string, message []byte) error
	Subscribe(topic string, handler EventHandler) error
}

type EventHandler func([]byte)
