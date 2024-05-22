# CHTTP Zap Logger

Package `chttp_logger_zap` with the implementation of the [middleware.Logger](https://github.com/spyzhov/chttp/blob/master/middleware/logger.go)
with the [zap.Logger](https://github.com/uber-go/zap) as the logger engine.

## Example

```go
package main

import (
	"context"

	"github.com/spyzhov/chttp"
	"github.com/spyzhov/chttp-logger-zap"
	"github.com/spyzhov/chttp/middleware"
	"go.uber.org/zap/zapcore"
)

func main() {
	client := chttp.NewJSON(nil)
	client.With(
		middleware.Trace(chttp_logger_zap.New("client", zapcore.InfoLevel, zapcore.ErrorLevel)),
		middleware.Debug(true, chttp_logger_zap.New("client", zapcore.DebugLevel, zapcore.DebugLevel)),
	)
	var result RandomUsers
	_ = client.GET(context.Background(), "https://randomuser.me/api/?results=10", nil, &result)
}

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
```

The example of the embedded version can be found in [logger_embed_test.go](logger_embed_test.go).

# License

MIT licensed. See the [LICENSE](LICENSE) file for details.
