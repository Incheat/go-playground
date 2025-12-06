package member

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/incheat/go-playground/services/user/internal/repository"
	"github.com/incheat/go-playground/services/user/pkg/model"
)

// ErrNotFound is returned when a requested record is not found.
var ErrMemberNotFound = errors.New("member not found")

// ErrMemberAlreadyExists is returned when a member already exists.
var ErrMemberAlreadyExists = errors.New("member already exists")

// Controller is the controller for the auth API.
type Controller struct {
	memberRepo MemberRepository
}

// AuthRepository is the interface for the auth repository.
type MemberRepository interface {
	GetMemberByEmail(ctx context.Context, email string) (*model.Member, error)
	CreateMember(ctx context.Context, email string, member *model.Member) error
}


// NewController creates a new Controller.
func NewController(memberRepo MemberRepository) *Controller {
	return &Controller{memberRepo: memberRepo}
}

// GetMemberByEmail gets a member by email.
func (c *Controller) GetMemberByEmail(ctx context.Context, email string) (*model.Member, error) {
	member, err := c.memberRepo.GetMemberByEmail(ctx, email)
	if err != nil && errors.Is(err, repository.ErrMemberNotFound) {
		return nil, ErrMemberNotFound
	}else if err != nil {
		return nil, err
	}
	return member, nil
}

// CreateMember creates a new member.
func (c *Controller) CreateMember(ctx context.Context, member *model.Member) error {
	member.ID = uuid.New().String()
	member.CreatedAt = time.Now()
	err := c.memberRepo.CreateMember(ctx, member.Email, member)
		if err != nil && errors.Is(err, repository.ErrMemberAlreadyExists) {
		return ErrMemberAlreadyExists
	}else if err != nil {
		return err
	}
	return nil
}