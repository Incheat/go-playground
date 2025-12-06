package http

import (
	"context"

	gen "github.com/incheat/go-playground/services/auth/internal/api/gen/oapi/public/server"
	"github.com/incheat/go-playground/services/auth/internal/constant"
	"github.com/incheat/go-playground/services/auth/internal/controller/auth"
)

// _ is a placeholder to ensure that Server implements the StrictServerInterface interface.
var _ gen.StrictServerInterface = (*Handler)(nil)

// Handler is the handler for the Auth API.
type Handler struct {
	ctrl *auth.Controller
}

// NewServer creates a new Server.
func NewHandler(ctrl *auth.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// Login is the handler for the Login endpoint.
func (h *Handler) Login(ctx context.Context, request gen.LoginRequestObject) (gen.LoginResponseObject, error) {
	email := string(request.Body.Email)
	password := request.Body.Password
	accessToken, refreshToken, err := h.ctrl.LoginWithEmailAndPassword(ctx, email, password)
	if err != nil {
		return gen.Login500JSONResponse{
			Error: err.Error(),
		}, err
	}

	return gen.Login200JSONResponse	{
		Body: gen.AuthResponse{
			AccessToken:  &accessToken,
			RefreshToken: &refreshToken,
		},
		Headers: gen.Login200ResponseHeaders{
			VersionId: constant.APIResponseVersionV1,
		},
	}, nil
}

// Logout is the handler for the Logout endpoint.
func (h *Handler) Logout(ctx context.Context, request gen.LogoutRequestObject) (gen.LogoutResponseObject, error) {
	return gen.Logout204Response{
		Headers: gen.Logout204ResponseHeaders{
			VersionId: constant.APIResponseVersionV1,
		},
	}, nil
}