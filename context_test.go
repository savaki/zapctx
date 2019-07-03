// Copyright 2019 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zapctx_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/savaki/zapctx"
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
			ctx    = zapctx.NewContext(parent, logger)
			want   = "boom"
		)

		zapctx.FromContext(ctx).Info(want)
		if !strings.Contains(got.String(), want) {
			t.Fatalf("got %v; wanted string to contain %v", got, want)
		}
	})

	t.Run("nop", func(t *testing.T) {
		ctx := context.Background()
		zapctx.FromContext(ctx).Info("boom")
		fmt.Printf("%#v\n", zapctx.FromContext(ctx))
	})
}
