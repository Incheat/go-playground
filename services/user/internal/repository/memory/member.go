package memory

import (
	"context"
	"sync"

	"github.com/incheat/go-playground/services/user/internal/repository"
	"github.com/incheat/go-playground/services/user/pkg/model"
)

// Repository defines a memory movie metadata repository.
type MemberRepository struct {
	sync.RWMutex
	data map[string]*model.Member
}

// NewRepository creates a new memory movie metadata repository.
func NewMemberRepository() *MemberRepository {
	return &MemberRepository{
		data: make(map[string]*model.Member),
	}
}

// GetMemberByEmail retrieves a member by email.
func (r *MemberRepository) GetMemberByEmail(_ context.Context, email string) (*model.Member, error) {
	r.RLock()
	defer r.RUnlock()
	member, ok := r.data[email]
	if !ok {
		return nil, repository.ErrMemberNotFound
	}
	return member, nil
}

// CreateMember creates a new member.
func (r *MemberRepository) CreateMember(_ context.Context, email string, member *model.Member) error {
	r.RLock()
	defer r.RUnlock()
	_, ok := r.data[email]
	if ok {
		return repository.ErrMemberAlreadyExists
	}

	r.Lock()
	defer r.Unlock()
	r.data[email] = member
	return nil
}