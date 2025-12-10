// Package authhandler defines the server for the Auth API.
package authhandler

import (
	"context"
	"fmt"

	servergen "github.com/incheat/go-playground/services/auth/internal/api/gen/oapi/public/server"
	"github.com/incheat/go-playground/services/auth/internal/constant"
	authservice "github.com/incheat/go-playground/services/auth/internal/service/auth"
)

// _ is a placeholder to ensure that Server implements the StrictServerInterface interface.
var _ servergen.StrictServerInterface = (*Server)(nil)

// Server is the server for the Auth API.
type Server struct {
	service *authservice.Service
}

// New creates a new Server.
func New(service *authservice.Service) *Server {
	return &Server{service: service}
}

// Login is the server for the Login endpoint.
func (h *Server) Login(ctx context.Context, request servergen.LoginRequestObject) (servergen.LoginResponseObject, error) {
	email := string(request.Body.Email)
	password := request.Body.Password
	userAgent := ""
	ipAddress := ""

	res, err := h.service.LoginWithEmailAndPassword(ctx, email, password, userAgent, ipAddress)
	if err != nil {
		return servergen.Login500JSONResponse{
			Error: err.Error(),
		}, err
	}

	accessToken := string(res.AccessToken)
	setCookie := fmt.Sprintf("refresh_token=%s; HttpOnly; Secure; SameSite=Lax; Path=/%s/%s; Max-Age=%d", res.RefreshToken, constant.APIResponseVersionV1, res.RefreshEndPoint, res.RefreshMaxAgeSec)

	return servergen.Login200JSONResponse{
		Body: servergen.AuthResponse{
			AccessToken: &accessToken,
		},
		Headers: servergen.Login200ResponseHeaders{
			VersionId: constant.APIResponseVersionV1,
			SetCookie: setCookie,
		},
	}, nil
}

// Logout is the server for the Logout endpoint.
func (h *Server) Logout(_ context.Context, _ servergen.LogoutRequestObject) (servergen.LogoutResponseObject, error) {
	return servergen.Logout204Response{
		Headers: servergen.Logout204ResponseHeaders{
			VersionId: constant.APIResponseVersionV1,
		},
	}, nil
}
