package authn

import (
	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

// Storage is storage for authentication
type Storage interface {
	GetUser(id string) (model.User, error)
	GetUserFromToken(token string) (model.User, error)
	GetHash(email string) (hash []byte, salt []byte, userID string, err error)
	AddToken(token string, userID string) error
	RemoveToken(token string) error
}

// Service is an authentication service
type Service struct {
	stor Storage
}

// New returns a new authentication service
func New(s Storage) (*Service, error) {
	return &Service{s}, nil
}
