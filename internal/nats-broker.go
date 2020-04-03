package internal

type natsBroker struct {

}

func newNatsBroker(config Config) Broker{
	return natsBroker{}
}

func (n natsBroker) Publish(topic string, message []byte) {

}

func (n natsBroker) Subscribe(topic string, handler EventHandler) {

}

