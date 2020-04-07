package main

import (
	"github.com/abdelhamidbakhta/senkyou/internal"
	"github.com/abdelhamidbakhta/senkyou/internal/log"
	"github.com/abdelhamidbakhta/senkyou/internal/net"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

func main() {
	config := internal.NewDefaultConfig()

	cmd := &cobra.Command{
		Use:   "senkyou",
		Short: "senkyou provides an Ethereum RPC gateway over message broker systems such as Kafka.",
		RunE:  run(&config),
	}
	cmd.PersistentFlags().StringVar(&config.BrokerType, "broker-type", config.BrokerType, "message broker type (nats, kafka)")
	cmd.PersistentFlags().StringVar(&config.KafkaUrl, "kafka-url", config.KafkaUrl, "kafka bootstrap server")
	cmd.PersistentFlags().StringVar(&config.NatsUrl, "nats-url", config.NatsUrl, "nats server url")
	cmd.PersistentFlags().BoolVar(&config.HttpEnabled, "http-enabled", config.HttpEnabled, "start http server for administration")
	cmd.PersistentFlags().IntVar(&config.HttpPort, "http-port", config.HttpPort, "http port")
	cmd.PersistentFlags().StringVar(&config.RpcUrl, "rpc-url", config.RpcUrl, "ethereum rpc url")
	cmd.PersistentFlags().StringVar(&config.TopicIncomingRpcRequests, "topic-rpc-requests", config.TopicIncomingRpcRequests, "topic to use for receiving incoming RPC requests")
	cmd.PersistentFlags().StringVar(&config.TopicOutgoingRpcResponses, "topic-rpc-responses", config.TopicOutgoingRpcResponses, "topic to use for pushing RPC responses")
	cmd.PersistentFlags().StringVar(&config.TopicErrors, "topic-errors", config.TopicErrors, "topic to use for error handling")
	cmd.PersistentFlags().Var(&config.LogLevel, "logging", "log level (DEBUG, INFO, WARN, ERROR)")
	err := cmd.Execute()
	logger := log.GetLoggerWithLevel(config.LogLevel.ZapLevel)
	defer logger.Sync()
	if err != nil {
		logger.Error("Failed to execute", zap.Error(err))
		os.Exit(1)
	}
}

func run(config *internal.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		broker, err := internal.NewBroker(*config)
		if err != nil {
			return err
		}
		senkyou, err := internal.NewSenkyou(*config, broker)
		if err != nil {
			return err
		}
		go senkyou.Start()
		if config.HttpEnabled {
			net.NewSenkyouServer(config.ListenAddr(), broker, config.LogLevel.ZapLevel).Start()
		}

		return nil
	}
}
