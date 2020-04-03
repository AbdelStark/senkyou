package internal

import "fmt"

type Config struct {
	KafkaUrl    string
	HttpEnabled bool
	HttpPort    int
}

func NewDefaultConfig() Config {
	return Config{
		KafkaUrl:    "127.0.0.1:9092",
		HttpEnabled: false,
		HttpPort:    8080,
	}
}

func (c Config) ListenAddr() string {
	return fmt.Sprintf(":%d", c.HttpPort)
}
