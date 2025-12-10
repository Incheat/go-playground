// Package logctx defines the context for the logger.
package logctx

import (
	"context"

	"go.uber.org/zap"
)

type loggerKey struct{}

// WithLogger adds a logger to the context.
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// From gets the logger from the context.
func From(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(loggerKey{}).(*zap.Logger)
	if !ok {
		// fallback to global logger
		return zap.L()
	}
	return logger
}
