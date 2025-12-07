// Package model defines the models for the auth service.
package model

import "time"

// AccessToken is a string that represents an access token.
type AccessToken string

// RefreshToken is a string that represents a refresh token.
type RefreshToken string

// RefreshTokenSession is a model for a refresh token session.
type RefreshTokenSession struct {
	ID        string
	MemberID  string
	TokenHash RefreshToken
	ExpiresAt time.Time
	CreatedAt time.Time
	RevokedAt time.Time
	UserAgent string
	IPAddress string
}
