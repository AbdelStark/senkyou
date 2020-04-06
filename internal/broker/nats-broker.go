package broker

import (
	"github.com/abdelhamidbakhta/senkyou/internal/log"
	"github.com/nats-io/nats.go"
)

var logger = log.ForceGetLogger()

type natsBroker struct {
	*nats.Conn
}

func NewNatsBroker(url string) (Broker, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return natsBroker{
		Conn: conn,
	}, nil
}

func (n natsBroker) Publish(topic string, message []byte) error {
	return n.Conn.Publish(topic, message)
}

func (n natsBroker) Subscribe(topic string, handler EventHandler) error {
	_, err := n.Conn.Subscribe(topic, func(msg *nats.Msg) {
		handler(msg.Data)
	})
	return err
}
