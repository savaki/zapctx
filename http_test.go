package zapctx

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/savaki/zapctx/zaputil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestMiddleware(t *testing.T) {
	var (
		buf        = bytes.NewBuffer(nil)
		config     = zap.NewDevelopmentEncoderConfig()
		core       = zapcore.NewCore(zapcore.NewConsoleEncoder(config), zaputil.NopSyncer(buf), zap.InfoLevel)
		logger     = zap.New(core)
		factory    = func(req *http.Request) *zap.Logger { return logger }
		middleware = Middleware(WithFactory(factory))
		text       = "hello world"
		target     = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			FromContext(req.Context()).Info(text)
		})
	)

	// When
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	handler := middleware(target)
	handler.ServeHTTP(w, req)

	// Then
	if got, want := w.Code, http.StatusOK; got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
	if got, want := buf.String(), text; !strings.Contains(got, want) {
		t.Fatalf("got %v; want substring %v", got, want)
	}
}

func TestMiddlewareNop(t *testing.T) {
	var (
		middleware = Middleware()
		target     = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			FromContext(req.Context()).Info("blah")
		})
	)

	// When
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	handler := middleware(target)
	handler.ServeHTTP(w, req)

	// Then
	if got, want := w.Code, http.StatusOK; got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
