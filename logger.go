package chttp_logger_zap

import (
	"context"
	"fmt"

	"github.com/spyzhov/chttp/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Logger *zap.Logger
	Info   zapcore.Level
	Error  zapcore.Level
}

func New(options ...Option) middleware.Logger {
	log := &Logger{
		Logger: zap.NewNop(),
		Info:   zapcore.DebugLevel,
		Error:  zapcore.ErrorLevel,
	}
	for _, o := range options {
		o(log)
	}
	return log
}

func (l *Logger) WithContext(_ context.Context) middleware.Logger {
	return &Logger{
		Error:  l.Error,
		Info:   l.Info,
		Logger: l.Logger,
	}
}

func (l *Logger) WithField(key string, value interface{}) middleware.Logger {
	lvl := l.Info
	if key == "error" {
		lvl = l.Error
	}
	return &Logger{
		Info:   lvl,
		Error:  l.Error,
		Logger: l.Logger.With(zap.Any(key, value)),
	}
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.Logger.Log(l.Info, fmt.Sprintf(format, args...))
}

func (l *Logger) WithLogger(log *zap.Logger) *Logger {
	return &Logger{
		Error:  l.Error,
		Info:   l.Info,
		Logger: log,
	}
}

func (l *Logger) WithError(lvl zapcore.Level) *Logger {
	return &Logger{
		Error:  lvl,
		Info:   l.Info,
		Logger: l.Logger,
	}
}

func (l *Logger) WithInfo(lvl zapcore.Level) *Logger {
	return &Logger{
		Error:  l.Error,
		Info:   lvl,
		Logger: l.Logger,
	}
}
