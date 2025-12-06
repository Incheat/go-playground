package repository

import "errors"

var (
	ErrMemberAlreadyExists = errors.New("member already exists")
	ErrMemberNotFound = errors.New("member not found")
)