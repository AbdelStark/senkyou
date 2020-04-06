package internal

import "github.com/abdelhamidbakhta/senkyou/internal/broker"

type Senkyou interface {
	Start()
}

func NewSenkyou(config Config, broker broker.Broker) (Senkyou, error) {
	return senkyou{
		Broker: broker,
	}, nil
}

type senkyou struct {
	broker.Broker
}

func (s senkyou) Start() {

}
