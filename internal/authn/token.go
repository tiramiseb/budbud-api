package authn

import (
	"crypto/rand"
	"fmt"

	authnerrors "github.com/tiramiseb/budbud-api/internal/authn/errors"
	"github.com/tiramiseb/budbud-api/internal/authn/scrypt"
	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

// GetUserFromToken returns the user linked to this token
func (s *Service) GetUserFromToken(token string) (model.User, error) {
	return s.stor.GetUserFromToken(token)
}

// CreateTokenFromCredentials creates an authentication token from a pair of credentials
func (s *Service) CreateTokenFromCredentials(email, password string) (string, model.User, error) {
	hash, salt, userID, err := s.stor.GetHash(email)
	if err != nil {
		return "", model.User{}, err
	}
	var ok bool
	// Debugging trick: set hash and salt of a user to "X" in order to accept all passwords
	if len(hash) == 1 && hash[0] == 'X' && len(salt) == 1 && salt[0] == 'X' {
		ok = true
	} else {
		ok, err = scrypt.Check(password, hash, salt)
	}
	if !ok {
		return "", model.User{}, authnerrors.ErrWrongEmailOrPassword
	}
	u, err := s.stor.GetUser(userID)
	if err != nil {
		return "", model.User{}, err
	}
	tokenB := make([]byte, 16)
	if _, err := rand.Read(tokenB); err != nil {
		return "", model.User{}, err
	}
	token := fmt.Sprintf("%x", tokenB)
	err = s.stor.AddToken(token, userID)
	if err != nil {
		return "", model.User{}, err
	}
	return token, u, nil

}

// RemoveToken removes a token
func (s *Service) RemoveToken(token string) error {
	return s.stor.RemoveToken(token)
}
