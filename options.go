package chttp_logger_zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(l *Logger)

func WithLogger(log *zap.Logger) Option {
	return func(l *Logger) {
		l.Logger = log
	}
}

func WithInfoLevel(lvl zapcore.Level) Option {
	return func(l *Logger) {
		l.Info = lvl
	}
}

func WithErrorLevel(lvl zapcore.Level) Option {
	return func(l *Logger) {
		l.Error = lvl
	}
}
