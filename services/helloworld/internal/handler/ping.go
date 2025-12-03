package handler

import (
	"context"

	gen "github.com/incheat/go-playground/services/helloworld/internal/api/gen/server"
)

var _ gen.StrictServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() *Server {
    return &Server{}
}

func (s *Server) PingV1(ctx context.Context, request gen.PingV1RequestObject) (gen.PingV1ResponseObject, error) {
	message := "pong"
	return gen.PingV1200JSONResponse{
		Body: gen.PingResponseV1{
			Message: &message,
		},
		Headers: gen.PingV1200ResponseHeaders{
			VersionId: "v1",
		},
	}, nil
}

func (s *Server) PingV2(ctx context.Context, request gen.PingV2RequestObject) (gen.PingV2ResponseObject, error) {
	message := "pong"
	return gen.PingV2200JSONResponse{
		Body: gen.PingResponseV2{
			Message: &message,
		},
		Headers: gen.PingV2200ResponseHeaders{
			VersionId: "v2",
		},
	}, nil
}