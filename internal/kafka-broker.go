package internal

type kafkaBroker struct {

}

func newKafkaBroker(config Config) Broker{
	return kafkaBroker{}
}

func (k kafkaBroker) Publish(topic string, message []byte) {

}

func (k kafkaBroker) Subscribe(topic string, handler EventHandler) {

}


