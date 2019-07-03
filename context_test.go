package zapx_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/savaki/zapx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type syncWriter struct {
	target io.Writer
}

func (s syncWriter) Write(p []byte) (n int, err error) {
	return s.target.Write(p)
}

func (s syncWriter) Sync() error {
	return nil
}

func TestContext(t *testing.T) {
	t.Run("ok", func(t *testing.T) {

		var (
			got    = bytes.NewBuffer(nil)
			w      = syncWriter{target: got}
			config = zap.NewDevelopmentEncoderConfig()
			core   = zapcore.NewCore(zapcore.NewConsoleEncoder(config), w, zapcore.InfoLevel)
			logger = zap.New(core)
			parent = context.Background()
			ctx    = zapx.NewContext(parent, logger)
			want   = "boom"
		)

		zapx.FromContext(ctx).Info(want)
		if !strings.Contains(got.String(), want) {
			t.Fatalf("got %v; wanted string to contain %v", got, want)
		}
	})

	t.Run("nop", func(t *testing.T) {
		ctx := context.Background()
		zapx.FromContext(ctx).Info("boom")
	})
}
