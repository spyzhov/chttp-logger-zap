package chttp_logger_zap

import (
	"context"

	"github.com/spyzhov/chttp"
	"github.com/spyzhov/chttp/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type RandomUsers struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
	} `json:"results"`
}

func ExampleLogger() {
	client := chttp.NewJSON(nil)
	client.With(
		middleware.Trace(New(WithLogger(zap.L()), WithInfoLevel(zapcore.InfoLevel))),
		middleware.Debug(true, New(WithLogger(zap.L()), WithErrorLevel(zapcore.DebugLevel))),
	)
	var result RandomUsers
	_ = client.GET(context.Background(), "https://randomuser.me/api/?results=10", nil, &result)
}
