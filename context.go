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

package zapctx

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
