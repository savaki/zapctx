package zaputil

import (
	"bytes"
	"io"
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestImplementsWriteSyncer(t *testing.T) {
	w := NopSyncer(bytes.NewBuffer(nil))
	v, ok := w.(zapcore.WriteSyncer)
	if !ok {
		t.Fatalf("expected zapcore.WriteSyncer got %T", v)
	}
}

func TestNop(t *testing.T) {
	var (
		buf  = bytes.NewBuffer(nil)
		w    = NopSyncer(buf)
		want = "hello world"
	)

	_, err := io.WriteString(w, want)
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}

	err = w.Sync()
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}
}
