// Package interceptor defines the interceptors for the user service.
package interceptor

import (
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

// DefaultChain returns a default chain of interceptors for the user service.
func DefaultChain(
	limiter *rate.Limiter,
) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		Recovery(),
		RateLimit(limiter),
		Logging(),
	}
}
