package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	db "github.com/incheat/go-playground/services/auth/internal/db/gen"
	"golang.org/x/crypto/bcrypt"
)

// Controller is the controller for the auth API.
type Controller struct {
	// q     *db.Queries
	repo AuthRepository
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
	CreateToken(userID string) (string, error)
	ParseUserID(tokenStr string) (string, error)
}

// AuthRepository is the interface for the auth repository.
type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	CreateUser(ctx context.Context, user db.CreateUserParams) error
}

// NewController creates a new Controller.
func NewController(repo AuthRepository, jwt JWTMaker, redis RedisClient) *Controller {
	return &Controller{repo: repo, jwt: jwt, redis: redis}
}

// Register registers a new user.
func (s *Controller) Register(ctx context.Context, email, password, name string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	id := uuid.NewString()
	return s.repo.CreateUser(ctx, db.CreateUserParams{
		ID:           id,
		Email:        email,
		PasswordHash: string(hash),
		Name:         name,
	})
}

// Login logs in a user.
func (s *Controller) Login(ctx context.Context, email, password string) (userID, accessToken, refreshToken string, err error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", "", errors.New("invalid credentials")
		}
		return "", "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", "", "", errors.New("invalid credentials")
	}

	access, err := s.jwt.CreateToken(u.ID)
	if err != nil {
		return "", "", "", err
	}

	refresh := uuid.NewString()
	// e.g. store refresh -> userID in Redis with longer TTL
	_ = s.redis.Set(ctx, "auth:refresh:"+refresh, u.ID, 7*24*time.Hour)

	return u.ID, access, refresh, nil
}