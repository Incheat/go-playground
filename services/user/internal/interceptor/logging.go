package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// Logging logs the request and response using Zap.
func Logging() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		start := time.Now()
		resp, err := handler(ctx, req)

		st, _ := status.FromError(err)
		log.Printf(
			"method=%s code=%s duration=%s",
			info.FullMethod,
			st.Code(),
			time.Since(start),
		)
		return resp, err
	}
}
