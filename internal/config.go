package internal

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	BrokerType  string
	NatsUrl string
	KafkaUrl    string
	HttpEnabled bool
	HttpPort    int
	RpcUrl      string
}

func NewDefaultConfig() Config {
	return Config{
		BrokerType:  Nats,
		NatsUrl:    "nats://127.0.0.1:4222",
		KafkaUrl:    "127.0.0.1:9092",
		HttpEnabled: false,
		HttpPort:    8080,
		RpcUrl:      "127.0.0.1:8545",
	}
}

func (c Config) string() string {
	payload, _ := json.MarshalIndent(c, "", "\t")
	return string(payload)
}

func (c Config) ListenAddr() string {
	return fmt.Sprintf(":%d", c.HttpPort)
}
