/*
Copyright 2024 Lucas de Ataides

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// The custom logger log levels
type LogLevel int

const (
	LOG_LEVEL_DEBUG LogLevel = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARNING
	LOG_LEVEL_ERROR
	LOG_LEVEL_FATAL
)

// Logger defines the logging structure for this application
type Logger struct {
	Log   zerolog.Logger
	Level LogLevel
}

// GetLogLevel transform a string in a LogLevel integer
func getLogLevel(logLevel string) LogLevel {
	switch strings.ToLower(logLevel) {
	case "debug":
		return LOG_LEVEL_DEBUG
	case "info":
		return LOG_LEVEL_INFO
	case "warning":
		return LOG_LEVEL_WARNING
	case "error":
		return LOG_LEVEL_ERROR
	case "fatal":
		return LOG_LEVEL_FATAL
	default:
		return LOG_LEVEL_INFO
	}
}

// InitializeLogger initializes the global logger.
func InitializeMainLogger(logFile string, logLevel string) *Logger {
	// Define output
	output, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// If file could not be open, use Stdout
		fmt.Printf("error loading %s. Defaulting to Stdout", logFile)
		output = os.Stdout
	}

	// Create logger
	log := zerolog.New(output).With().Timestamp().Logger()
	return &Logger{
		Log:   log,
		Level: getLogLevel(logFile),
	}
}

// Log in Debug level
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.Level <= LOG_LEVEL_DEBUG {
		msg := fmt.Sprintf(format, v...)
		l.Log.Debug().Msg(msg)
	}
}

// Log in Info level
func (l *Logger) Info(format string, v ...interface{}) {
	if l.Level <= LOG_LEVEL_INFO {
		msg := fmt.Sprintf(format, v...)
		l.Log.Info().Msg(msg)
	}
}

// Log in Warn level
func (l *Logger) Warn(format string, v ...interface{}) {
	if l.Level <= LOG_LEVEL_WARNING {
		msg := fmt.Sprintf(format, v...)
		l.Log.Warn().Msg(msg)
	}
}

// Log in Error level
func (l *Logger) Error(format string, v ...interface{}) {
	if l.Level <= LOG_LEVEL_ERROR {
		msg := fmt.Sprintf(format, v...)
		l.Log.Error().Msg(msg)
	}
}

// Log in Fatal level
func (l *Logger) Fatal(format string, v ...interface{}) {
	if l.Level <= LOG_LEVEL_FATAL {
		msg := fmt.Sprintf(format, v...)
		l.Log.Fatal().Msg(msg)
	}
}
