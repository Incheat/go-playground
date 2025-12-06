package memory

import (
	"context"
	"sync"

	"github.com/incheat/go-playground/services/auth/internal/repository"
	"github.com/incheat/go-playground/services/auth/pkg/model"
)

// Repository defines a memory movie metadata repository.
type RefreshTokenRepository struct {
	sync.RWMutex
	data map[string]*model.RefreshToken
}

// NewRepository creates a new memory movie metadata repository.
func NewRefreshTokenRepository() *RefreshTokenRepository {
	return &RefreshTokenRepository{
		data: make(map[string]*model.RefreshToken),
	}
}

// GetRefreshTokenByID retrieves a refresh token by ID.
func (r *RefreshTokenRepository) GetRefreshTokenByID(_ context.Context, id string) (*model.RefreshToken, error) {
	r.RLock()
	defer r.RUnlock()
	refreshToken, ok := r.data[id]
	if !ok {
		return nil, repository.ErrRefreshTokenNotFound
	}
	return refreshToken, nil
}

// CreateRefreshToken creates a new refresh token.
func (r *RefreshTokenRepository) CreateRefreshToken(_ context.Context, id string, refreshToken *model.RefreshToken) error {
	r.RLock()
	defer r.RUnlock()
	_, ok := r.data[id]
	if ok {
		return repository.ErrRefreshTokenAlreadyExists
	}

	r.Lock()
	defer r.Unlock()
	r.data[id] = refreshToken
	return nil
}