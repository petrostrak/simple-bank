package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken create a new token for a specific username and duration
	CreateToken(string, time.Duration) (string, error)

	// VerifyToken checks if the token is valid
	VerifyToken(string) (*Payload, error)
}
