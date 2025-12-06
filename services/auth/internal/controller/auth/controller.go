package auth

import (
	"context"
	"time"

	"github.com/incheat/go-playground/services/auth/pkg/model"
)

// Controller is the controller for the auth API.
type Controller struct {
	refreshTokenRepo RefreshTokenRepository
	jwt JWTMaker
	redis RedisClient
}

// RedisClient is the interface for the Redis client.
type RedisClient interface {
	Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error
	Del(ctx context.Context, keys ...string) error
}

// JWTMaker is the interface for the JWT maker.
type JWTMaker interface {
	CreateToken(ID string) (model.AccessToken, error)
	ParseID(token string) (string, error)
}

// RefreshTokenRepository is the interface for the refresh token repository.
type RefreshTokenRepository interface {
	GetRefreshTokenByID(ctx context.Context, id string) (*model.RefreshToken, error)
	CreateRefreshToken(ctx context.Context, id string, refreshToken *model.RefreshToken) error
}

// NewController creates a new Controller.
func NewController(refreshTokenRepo RefreshTokenRepository, jwt JWTMaker, redis RedisClient) *Controller {
	return &Controller{refreshTokenRepo: refreshTokenRepo, jwt: jwt, redis: redis}
}

// LoginWithEmailAndPassword logs in a user with email and password.
func (c *Controller) LoginWithEmailAndPassword(ctx context.Context, email string, password string) (string, string, error) {
	accessToken, err := c.jwt.CreateToken(email)
	if err != nil {
		return "", "", err
	}
	return string(accessToken), "", nil
}