package logger

import (
	"time"

	"gorm.io/gorm/logger"
)

type Option func(*Logger)

func WithPath(path string) Option {
	return func(l *Logger) {
		l.path = path
	}
}

func WithSlowThreshold(slowThreshold time.Duration) Option {
	return func(l *Logger) {
		l.slowThreshold = slowThreshold
	}
}

func WithLogLevel(logLevel logger.LogLevel) Option {
	return func(l *Logger) {
		l.logLevel = logLevel
	}
}

func WithMaxSize(maxSize int) Option {
	return func(l *Logger) {
		l.maxSize = maxSize
	}
}

func WithMaxBackups(maxBackups int) Option {
	return func(l *Logger) {
		l.maxBackups = maxBackups
	}
}

func WithMaxAge(maxAge int) Option {
	return func(l *Logger) {
		l.maxAge = maxAge
	}
}
