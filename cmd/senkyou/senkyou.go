package main

import (
	"github.com/abdelhamidbakhta/senkyou/internal"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

func main() {
	config := internal.NewDefaultConfig()
	logger, _ := zap.NewDevelopment()
	cmd := &cobra.Command{
		Use:   "senkyou",
		Short: "senkyou provides an Ethereum RPC gateway over message broker systems such as Kafka.",
		RunE:  run(&config),
	}
	cmd.PersistentFlags().StringVar(&config.KafkaUrl, "kafka-url", config.KafkaUrl, "kafka bootstrap server")
	cmd.PersistentFlags().BoolVar(&config.HttpEnabled, "http-enabled", config.HttpEnabled, "start http server for administration")
	cmd.PersistentFlags().IntVar(&config.HttpPort, "http-port", config.HttpPort, "http port")
	cmd.PersistentFlags().StringVar(&config.RpcUrl, "rpc-url", config.RpcUrl, "ethereum rpc url")

	err := cmd.Execute()
	if err != nil {
		logger.Error("Failed to execute", zap.Error(err))
		os.Exit(1)
	}
}

func run(config *internal.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if config.HttpEnabled {
			internal.NewSenkyouServer(*config).Start()
		}
		senkyou, err := internal.NewSenkyou(*config)
		if err != nil {
			return err
		}
		senkyou.Start()
		return nil
	}
}
