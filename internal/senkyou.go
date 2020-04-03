package internal

type Senkyou interface {
	Start()
}

func NewSenkyou(config Config) (Senkyou, error) {
	broker, err := NewBroker(config)
	if err != nil {
		return nil, err
	}
	return senkyou{
		Broker: broker,
	}, nil
}

type senkyou struct {
	Broker
}

func (s senkyou) Start() {

}
