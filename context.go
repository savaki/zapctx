package zapx

import (
	"context"

	"go.uber.org/zap"
)

type key int

// contextKey holds our zap logger
const contextKey key = 1

// nop logger to ensure FromContext always returns something
var nop = zap.NewNop()

// FromContext retrieves the logger from the given Context.  If no logger was found,
// FromContext returns the nop logger
func FromContext(ctx context.Context) *zap.Logger {
	v := ctx.Value(contextKey)
	logger, ok := v.(*zap.Logger)
	if !ok {
		return nop
	}
	return logger
}

// NewContext returns a new context with the given logger attached
func NewContext(parent context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(parent, contextKey, logger)
}
