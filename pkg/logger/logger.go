package logger

import (
	"gojob/config"
	"os"

	"github.com/rs/zerolog"
)

// Ver 1
//func NewLogger(config *config.LogConfig) *zerolog.Logger {
//	zerolog.SetGlobalLevel(zerolog.Level(config.Level))
//	var logger zerolog.Logger
//	if config.Format == "json" {
//		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
//	} else {
//		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
//		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
//
//	}
//	return &logger
//
//}

func NewLogger(config *config.LogConfig) *zerolog.Logger {
	var logger zerolog.Logger

	if config.Format == "json" {
		logger = zerolog.New(os.Stdout)
	} else {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02 15:04:05",
		}
		logger = zerolog.New(consoleWriter)
	}

	logger = logger.With().
		Timestamp().
		Str("service", "gojob").
		Logger()

	level := zerolog.Level(config.Level)
	logger = logger.Level(level)

	return &logger
}
