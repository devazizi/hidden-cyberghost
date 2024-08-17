package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Logger struct {
	Logger *zerolog.Logger
}

func NewLogger() *Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	logger := zerolog.New(output).With().Timestamp().Logger()

	return &Logger{
		Logger: &logger,
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.Logger.Debug().Msgf(format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.Logger.Info().Msgf(format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.Logger.Error().Msgf(format, args...)
}
