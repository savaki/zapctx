package zapctx

import (
	"net/http"

	"go.uber.org/zap"
)

type options struct {
	makeLogger func(req *http.Request) *zap.Logger
}

// Option provides a functional
type Option func(*options)

// WithFactory allows a custom logger to be returned
func WithFactory(fn func(req *http.Request) *zap.Logger) Option {
	return func(o *options) {
		o.makeLogger = fn
	}
}

// InjectLogger returns handler to inject logger into request context
func InjectLogger(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			req = req.WithContext(NewContext(req.Context(), logger))
			handler.ServeHTTP(w, req)
		})
	}
}

// Middleware provides web middleware to inject a logger into the request context
func Middleware(opts ...Option) func(h http.Handler) http.Handler {
	nop := zap.NewNop()
	options := options{
		makeLogger: func(req *http.Request) *zap.Logger {
			return nop
		},
	}

	for _, opt := range opts {
		opt(&options)
	}

	return func(original http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			logger := options.makeLogger(req)
			ctx := NewContext(req.Context(), logger)
			req = req.WithContext(ctx)
			original.ServeHTTP(w, req)
		})
	}
}
