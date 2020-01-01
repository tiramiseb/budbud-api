package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	authnerrors "github.com/tiramiseb/budbud-api/internal/authn/errors"
	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

// GetUserFromToken returns the user from a token
func (s *Service) GetUserFromToken(token string) (model.User, error) {
	// This is currently dumb because we only return the email, which is the ID. However, as soon as there are other parameters in the table, it will make sense
	user := model.User{}
	if err := s.db.QueryRow("SELECT email, email FROM user INNER JOIN token ON user.email = token.user_email WHERE token.token=?", token).Scan(&user.ID, &user.Email); err != nil {
		return user, fmt.Errorf("Cannot get user from token: %w", err)
	}
	return user, nil
}

// GetHash returns a hash, a salt and a user ID from an email address
func (s *Service) GetHash(email string) (hash []byte, salt []byte, userID string, err error) {
	userID = email
	err = s.db.QueryRow("SELECT passhash, passsalt FROM user WHERE email=?;", email).Scan(&hash, &salt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, "", authnerrors.ErrWrongEmailOrPassword
		}
		return nil, nil, "", fmt.Errorf("Cannot get hash: %w", err)
	}
	return
}

// AddToken stores a token and the ID of its user
func (s *Service) AddToken(token string, userID string) error {
	if _, err := s.db.Exec("INSERT INTO token (token, user_email) VALUES (?, ?);", token, userID); err != nil {
		return fmt.Errorf("Cannot add token: %w", err)
	}
	return nil
}

// RemoveToken removes a token
func (s *Service) RemoveToken(token string) error {
	if _, err := s.db.Exec("DELETE FROM token WHERE token=?", token); err != nil {
		return fmt.Errorf("Cannot remove token: %w", err)
	}
	return nil
}
