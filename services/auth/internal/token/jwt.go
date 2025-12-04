package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Maker is a JWT maker.
type Maker struct {
	secret []byte
	expire time.Duration
}

// NewMaker creates a new JWT maker.
func NewMaker(secret string, minutes int) *Maker {
	return &Maker{
		secret: []byte(secret),
		expire: time.Duration(minutes) * time.Minute,
	}
}

// CreateToken creates a new JWT token for a user.
func (m *Maker) CreateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(m.expire).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(m.secret)
}

// ParseUserID parses the user ID from a JWT token.
// Returns the user ID if the token is valid, otherwise returns an error.
func (m *Maker) ParseUserID(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return m.secret, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrInvalidKey
	}
	sub, _ := claims["sub"].(string)
	return sub, nil
}
