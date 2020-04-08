package main

import (
	"github.com/abdelhamidbakhta/senkyou/internal"
	"github.com/abdelhamidbakhta/senkyou/internal/config"
	"github.com/abdelhamidbakhta/senkyou/internal/log"
	"github.com/abdelhamidbakhta/senkyou/internal/net"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"go.elastic.co/apm/module/apmot"
	"go.uber.org/zap"
	"os"
)

func main() {
	senkyouConfig := config.NewDefaultConfig()

	cmd := &cobra.Command{
		Use:   "senkyou",
		Short: "senkyou provides an Ethereum RPC gateway over message broker systems such as BrokerKafka.",
		RunE:  run(&senkyouConfig),
	}
	cmd.PersistentFlags().StringVar(&senkyouConfig.BrokerType, "broker-type", senkyouConfig.BrokerType, "message broker type (nats, kafka)")
	cmd.PersistentFlags().StringVar(&senkyouConfig.KafkaUrl, "kafka-url", senkyouConfig.KafkaUrl, "kafka bootstrap server")
	cmd.PersistentFlags().StringVar(&senkyouConfig.NatsUrl, "nats-url", senkyouConfig.NatsUrl, "nats server url")
	cmd.PersistentFlags().BoolVar(&senkyouConfig.HttpEnabled, "http-enabled", senkyouConfig.HttpEnabled, "start http server for administration")
	cmd.PersistentFlags().IntVar(&senkyouConfig.HttpPort, "http-port", senkyouConfig.HttpPort, "http port")
	cmd.PersistentFlags().StringVar(&senkyouConfig.RpcUrl, "rpc-url", senkyouConfig.RpcUrl, "ethereum rpc url")
	cmd.PersistentFlags().StringVar(&senkyouConfig.TopicIncomingRpcRequests, "topic-rpc-requests", senkyouConfig.TopicIncomingRpcRequests, "topic to use for receiving incoming RPC requests")
	cmd.PersistentFlags().StringVar(&senkyouConfig.TopicOutgoingRpcResponses, "topic-rpc-responses", senkyouConfig.TopicOutgoingRpcResponses, "topic to use for pushing RPC responses")
	cmd.PersistentFlags().StringVar(&senkyouConfig.TopicErrors, "topic-errors", senkyouConfig.TopicErrors, "topic to use for error handling")
	cmd.PersistentFlags().Var(&senkyouConfig.LogLevel, "logging", "log level (DEBUG, INFO, WARN, ERROR)")
	cmd.PersistentFlags().BoolVar(&senkyouConfig.ApmEnabled, "apm-enabled", senkyouConfig.ApmEnabled, "enable application performance monitoring using elk stack")

	err := cmd.Execute()
	logger := log.GetLogger(senkyouConfig)
	defer logger.Sync()
	if err != nil {
		logger.Error("Failed to execute", zap.Error(err))
		os.Exit(1)
	}
}

func run(cfg *config.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if cfg.ApmEnabled{
			opentracing.SetGlobalTracer(apmot.New())
		}
		broker, err := internal.NewBroker(*cfg)
		if err != nil {
			return err
		}
		senkyou, err := internal.NewSenkyou(*cfg, broker)
		if err != nil {
			return err
		}
		go senkyou.Start()
		if cfg.HttpEnabled {
			net.NewSenkyouServer(*cfg, broker, cfg.LogLevel.ZapLevel).Start()
		}

		return nil
	}
}
