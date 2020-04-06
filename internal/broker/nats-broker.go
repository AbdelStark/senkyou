package broker

type natsBroker struct {
	url string
}

func NewNatsBroker(url string) Broker {
	return natsBroker{
		url: url,
	}
}

func (n natsBroker) Publish(topic string, message []byte) {

}

func (n natsBroker) Subscribe(topic string, handler EventHandler) {

}
