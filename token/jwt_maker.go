package token

import (
	"fmt"
	"time"
)

const (
	minSecretKeySie = 32
)

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySie {
		return nil, fmt.Errorf("invalid key size: secret key must be at least %d characters", minSecretKeySie)
	}

	return &JWTMaker{secretKey}, nil
}

func (m *JWTMaker) CreateToken(username string, duration time.Duration) (string, error)

func (m *JWTMaker) VerifyToken(token string) (*Payload, error)
