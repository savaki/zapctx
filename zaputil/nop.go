package zaputil

import (
	"io"

	"go.uber.org/zap/zapcore"
)

type nopSyncer struct {
	target io.Writer
}

func (nop *nopSyncer) Write(p []byte) (n int, err error) {
	return nop.target.Write(p)
}

func (nop *nopSyncer) Sync() error {
	return nil
}

// NopSyncer allows any writer to
func NopSyncer(w io.Writer) zapcore.WriteSyncer {
	return &nopSyncer{target: w}
}
