package log

import "go.uber.org/zap"

func GetLogger() (*zap.Logger, error) {
	//return zap.NewProduction()
	return zap.NewDevelopment()
}

func ForceGetLogger() *zap.Logger {
	logger, _ := GetLogger()
	return logger
}
