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
		Run:   run(&config),
	}
	cmd.PersistentFlags().StringVar(&config.KafkaUrl, "kafkaUrl", config.KafkaUrl, "kafka bootstrap server (default is 127.0.0.1:9092)")
	cmd.PersistentFlags().BoolVar(&config.HttpEnabled, "http-enabled", config.HttpEnabled, "start http server for administration")
	cmd.PersistentFlags().IntVar(&config.HttpPort, "http-port", config.HttpPort, "kafka bootstrap server (default is 127.0.0.1:9092)")

	err := cmd.Execute()
	if err != nil {
		logger.Error("Failed to execute", zap.Error(err))
		os.Exit(1)
	}
}

func run(config *internal.Config) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if config.HttpEnabled {
			internal.NewSenkyouServer(*config).Start()
		}
	}
}
