package broker

type kafkaBroker struct {
	url string
}

func NewKafkaBroker(url string) Broker {
	return kafkaBroker{
		url: url,
	}
}

func (k kafkaBroker) Publish(topic string, message []byte) {

}

func (k kafkaBroker) Subscribe(topic string, handler EventHandler) {

}
