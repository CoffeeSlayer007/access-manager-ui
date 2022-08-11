package logger

import (
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once

/*
Format is the contract of the Log. Use data field to populate dynamic fields.
*/
type Format struct {
	Event    string
	Message  string
	Endpoint string
	Data     map[string]string
}

/*
IRapidoLogger is an interface for the logger. You could use this for mocking the logger if needed.
*/
type IRapidoLogger interface {
	Debug(log Format)
	Info(log Format)
	Error(log Format)
	Warn(log Format)
}

/*
RapidoLogger is a wrapper struct over Zap
*/
type RapidoLogger struct {
	*zap.Logger
}

func appendIfNonEmpty(key string, value string, result []zap.Field) []zap.Field {
	if value != "" {
		result = append(result, zap.String(key, value))
	}
	return result
}

func getZapFields(log Format) []zap.Field {

	result := []zap.Field{}

	result = appendIfNonEmpty("event", log.Event, result)
	result = appendIfNonEmpty("endpoint", log.Endpoint, result)
	result = appendIfNonEmpty("message", log.Message, result)
	result = appendIfNonEmpty("timestamp", time.Now().Format(time.RFC3339), result)

	if log.Data != nil {
		result = append(result, zap.Any("data", log.Data))
	}

	return result
}

/*
Debug log
*/
func Debug(log Format) {
	loggerWithNs.Debug(log.Message, getZapFields(log)...)
}

/*
Info log
*/
func Info(log Format) {
	loggerWithNs.Info(log.Message, getZapFields(log)...)
}

/*
Warn log
*/
func Warn(log Format) {
	loggerWithNs.Warn(log.Message, getZapFields(log)...)
}

/*
Error log
*/
func Error(log Format) {
	loggerWithNs.Error(log.Message, getZapFields(log)...)
}

var loggerWithNs *zap.Logger

/*
Init initializes a logger and returns it. As a practice, call this from your init stage and pass around the logger.
*/
func Init(logLevel string) {

	once.Do(func() {
		config := zap.Config{
			Encoding:         "json",
			Level:            getLogLevel(logLevel),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{

				LevelKey:    "level",
				EncodeLevel: zapcore.LowercaseLevelEncoder,
			},
		}

		logger, _ := config.Build()
		loggerWithNs = logger.With(zap.Namespace("message"))

		Info(Format{Message: "Logger successfully initialized"})
		defer logger.Sync()
		defer loggerWithNs.Sync()
	})
}

func getLogLevel(level string) zap.AtomicLevel {
	switch strings.ToLower(level) {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}
