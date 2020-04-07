package log

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	Debug = "DEBUG"
	Info  = "INFO"
	Warn  = "WARN"
	Error = "ERROR"
)

func GetLoggerWithLevel(level zapcore.Level) *zap.Logger {
	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	atom.SetLevel(level)
	return logger
}

func NewLogLevelFlag() LogLevelFlag {
	return LogLevelFlag{
		LevelString: "DEBUG",
		ZapLevel:    zapcore.DebugLevel,
	}
}

type LogLevelFlag struct {
	LevelString string
	ZapLevel    zapcore.Level
}

func (l LogLevelFlag) String() string {
	return l.LevelString
}

func (l *LogLevelFlag) Set(str string) error {
	levelString := strings.ToUpper(str)
	switch levelString {
	case Debug:
		l.ZapLevel = zapcore.DebugLevel
		return nil
	case Info:
		l.ZapLevel = zapcore.InfoLevel
		return nil
	case Warn:
		l.ZapLevel = zapcore.WarnLevel
		return nil
	case Error:
		l.ZapLevel = zapcore.ErrorLevel
		return nil
	}
	return errors.New(fmt.Sprintf("unknown log level: %s", str))
}

func (l LogLevelFlag) Type() string {
	return "logLevel"
}
