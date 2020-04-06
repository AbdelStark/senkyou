package broker

type kafkaBroker struct {
	url string
}

func NewKafkaBroker(url string) (Broker, error) {
	return kafkaBroker{
		url: url,
	}, nil
}

func (k kafkaBroker) Publish(topic string, message []byte) error {
	return nil
}

func (k kafkaBroker) Subscribe(topic string, handler EventHandler) error {
	return nil
}
