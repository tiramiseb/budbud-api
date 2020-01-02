package sqlite

import (
	"fmt"

	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

// GetUser returns a user from its ID
func (s *Service) GetUser(id string) (model.User, error) {
	// user ID is email
	// This is currently dumb because we only return the email, which is the ID. However, as soon as there are other parameters in the table, it will make sense
	user := model.User{}
	if err := s.db.QueryRow("SELECT email, email FROM user WHERE email=?", id).Scan(&user.ID, &user.Email); err != nil {
		return user, fmt.Errorf("Cannot get user data: %w", err)
	}
	return user, nil
}
