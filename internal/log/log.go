/*
  Filename: log.go
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description: Configures the structured logger for the application.
*/

package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger is a package-level variable that can be accessed from other parts of the application.
var Logger zerolog.Logger

// InitLogger sets up the logger with the specified level.
func InitLogger(level string) {
	var logLevel zerolog.Level

	switch level {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "verbose":
		logLevel = zerolog.InfoLevel
	case "quiet":
		logLevel = zerolog.Disabled
	default:
		logLevel = zerolog.ErrorLevel
	}

	Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(logLevel).
		With().
		Timestamp().
		Logger()
}
