// Package log provides a logging interface to send logs from plugins to Grafana server.
package log

import (
	"github.com/hashicorp/go-hclog"
)

// defaultLogger is the default logger.
// This can be overridden with SetLogger
var defaultLogger = New()

// Logger is the main Logger interface.
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// New creates a new logger.
func New() Logger {
	return &hclogWrapper{
		logger: hclog.New(&hclog.LoggerOptions{
			// Use debug as level since anything less severe is suppressed.
			Level: hclog.Debug,
			// Use JSON format to make the output in Grafana format and work
			// when using multiple arguments such as Debug("message", "key", "value").
			JSONFormat: true,
		}),
	}
}

type hclogWrapper struct {
	logger hclog.Logger
}

func (l *hclogWrapper) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, args...)
}

func (l *hclogWrapper) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

func (l *hclogWrapper) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, args...)
}

func (l *hclogWrapper) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
}

func DefaultLogger() Logger {
	return defaultLogger
}

// SetLogger is used by consumers to override the DefaultLogger
func SetLogger(logger Logger) {
	defaultLogger = logger
}
