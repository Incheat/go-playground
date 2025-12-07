// Package handler defines the handlers for the Auth API.
package handler

import (
	"context"
	"fmt"

	gen "github.com/incheat/go-playground/services/auth/internal/api/gen/oapi/public/server"
	"github.com/incheat/go-playground/services/auth/internal/constant"
	"github.com/incheat/go-playground/services/auth/internal/controller/auth"
	ginmiddleware "github.com/oapi-codegen/gin-middleware"
)

// _ is a placeholder to ensure that Server implements the StrictServerInterface interface.
var _ gen.StrictServerInterface = (*Handler)(nil)

// Handler is the handler for the Auth API.
type Handler struct {
	ctrl *auth.Controller
}

// NewHandler creates a new Handler.
func NewHandler(ctrl *auth.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// Login is the handler for the Login endpoint.
func (h *Handler) Login(ctx context.Context, request gen.LoginRequestObject) (gen.LoginResponseObject, error) {
	email := string(request.Body.Email)
	password := request.Body.Password
	gc := ginmiddleware.GetGinContext(ctx)
	if gc == nil {
		return gen.Login500JSONResponse{
			Error: "gin context not found",
		}, nil
	}
	userAgent := gc.Request.UserAgent()
	ipAddress := gc.ClientIP()

	res, err := h.ctrl.LoginWithEmailAndPassword(ctx, email, password, userAgent, ipAddress)
	if err != nil {
		return gen.Login500JSONResponse{
			Error: err.Error(),
		}, err
	}

	accessToken := string(res.AccessToken)
	setCookie := fmt.Sprintf("refresh_token=%s; HttpOnly; Secure; SameSite=Lax; Path=/%s/%s; Max-Age=%d", res.RefreshToken, constant.APIResponseVersionV1, res.RefreshEndPoint, res.RefreshMaxAgeSec)

	return gen.Login200JSONResponse{
		Body: gen.AuthResponse{
			AccessToken: &accessToken,
		},
		Headers: gen.Login200ResponseHeaders{
			VersionId: constant.APIResponseVersionV1,
			SetCookie: setCookie,
		},
	}, nil
}

// Logout is the handler for the Logout endpoint.
func (h *Handler) Logout(_ context.Context, _ gen.LogoutRequestObject) (gen.LogoutResponseObject, error) {
	return gen.Logout204Response{
		Headers: gen.Logout204ResponseHeaders{
			VersionId: constant.APIResponseVersionV1,
		},
	}, nil
}
