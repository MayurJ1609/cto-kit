// Package token provides a set of interfaces for service
// authorization through [JSON Web Tokens](https://jwt.io/).
package token

import (
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	// DefaultTimeout a set for the "exp" (expiration time) claim
	// identifies the expiration time on or after which the
	// JWT MUST NOT be accepted for processing
	defaultTimeout = 30 * time.Minute
)

// Claims holds token claims
type Claims struct {
	Phone    string   `json:"phone"`
	UserID   string   `json:"userID"`
	DeviceID string   `json:"deviceID"`
	Scope    []string `json:"scope,omitempty"`
}

type claim struct {
	Phone    string   `json:"identity"`
	UserID   string   `json:"id"`
	DeviceID string   `json:"device"`
	Scope    []string `json:"scope,omitempty"`
	jwt.RegisteredClaims
}

// Sign can be used for signing or verifying tokens.
type Sign interface {
	Verify(ctx context.Context, token string) (Claims, error)
	Generate(ctx context.Context, claims Claims, opts ...Option) (string, error)
}

type sign struct {
	algorithm *jwt.SigningMethodHMAC
	secret    string
}

// New returns JWT sign instance
func New(secret string) Sign {
	sign := &sign{
		algorithm: jwt.SigningMethodHS256,
	}
	return sign
}

// Generate the JWT signed token
func (sign *sign) Generate(ctx context.Context, claims Claims, opts ...Option) (string, error) {
	option := &option{
		expire: defaultTimeout,
	}
	for _, opt := range opts {
		opt(option)
	}

	expiresAt := jwt.NewNumericDate(time.Now().Add(option.expire))
	claim := claim{
		UserID:   claims.UserID,
		DeviceID: claims.DeviceID,
		Scope:    claims.Scope,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
		},
	}
	return jwt.NewWithClaims(
		sign.algorithm,
		&claim,
	).SignedString([]byte(sign.secret))
}

// Verify verify the given access token
func (sign *sign) Verify(ctx context.Context, token string) (Claims, error) {

	token = strings.ReplaceAll(token, "Bearer ", "")
	claim := claim{}
	t, err := jwt.ParseWithClaims(
		token,
		&claim,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(sign.secret), nil
		},
	)
	if err != nil || !t.Valid {
		return Claims{}, ErrInvalidToken
	}
	return Claims{
		UserID:   claim.UserID,
		DeviceID: claim.DeviceID,
		Scope:    claim.Scope,
	}, nil
}
