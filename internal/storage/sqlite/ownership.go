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

// GetWorkspacesForUserID returns the list of workspaces for a user
func (s *Service) GetWorkspacesForUserID(userID string) ([]*model.Workspace, error) {
	rows, err := s.db.Query("SELECT id FROM workspace INNER JOIN workspace_user ON workspace_user.workspace_id = workspace.id INNER JOIN user ON workspace_user.user_email = user.email")
	if err != nil {
		return nil, fmt.Errorf("Cannot get workspaces for user ID %s: %w", userID, err)
	}
	defer rows.Close()
	var workspaces []*model.Workspace
	for rows.Next() {
		var workspace *model.Workspace
		if err := rows.Scan(&workspace.ID); err != nil {
			return nil, err
		}
		workspaces = append(workspaces, workspace)
	}
	return workspaces, nil
}
