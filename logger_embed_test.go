package chttp_logger_zap

import (
	"context"

	"github.com/spyzhov/chttp/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type EmbeddedLogger struct {
	Logger
}

var _ middleware.Logger = (*EmbeddedLogger)(nil)

func NewEmbeddedLogger(client string) *EmbeddedLogger {
	return &EmbeddedLogger{
		Logger: *New(
			WithLogger(zap.L().With(zap.String("client", client))),
			WithInfoLevel(zapcore.InfoLevel),
		).(*Logger),
	}
}

func (l *EmbeddedLogger) WithContext(ctx context.Context) middleware.Logger {
	return &EmbeddedLogger{
		Logger: *l.Logger.WithLogger(l.Logger.Logger.With(zap.Any("request_id", ctx.Value("X-Request-ID")))),
	}
}

func (l *EmbeddedLogger) WithField(key string, value interface{}) middleware.Logger {
	if key == "request_time" { // replacement
		key = "duration"
	}
	return &EmbeddedLogger{
		Logger: *l.Logger.WithField(key, value).(*Logger),
	}
}
