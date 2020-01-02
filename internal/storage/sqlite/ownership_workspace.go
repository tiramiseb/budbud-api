package sqlite

import (
	"fmt"

	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

// AddWorkspace adds a new workspace and puts the given user as its owner
func (s *Service) AddWorkspace(userID, name string) (*model.Workspace, error) {
	// if _, err := s.db.Exec("INSERT INTO workspace (id, name) VALUES")
	// TODO
	return nil, nil
}

// GetWorkspacesOwned returns the list of workspaces owned by a user
func (s *Service) GetWorkspacesOwned(userID string) ([]*model.Workspace, error) {
	rows, err := s.db.Query(
		`SELECT workspace.id, workspace.name
		FROM workspace
		WHERE workspace.owner_email=?`,
		userID)
	if err != nil {
		return nil, fmt.Errorf("Cannot get workspaces owned by user ID %s: %w", userID, err)
	}
	defer rows.Close()
	var workspaces []*model.Workspace
	for rows.Next() {
		workspace := model.Workspace{}
		if err := rows.Scan(&workspace.ID, &workspace.Name); err != nil {
			return nil, err
		}
		workspaces = append(workspaces, &workspace)
	}
	return workspaces, nil
}

// GetWorkspacesGuest returns the list of workspaces a user is invited to
func (s *Service) GetWorkspacesGuest(userID string) ([]*model.Workspace, error) {
	rows, err := s.db.Query(
		`SELECT workspace.id, workspace.name, user.email, user.email
		FROM workspace
		INNER JOIN user ON workspace.owner_email=user.email
		INNER JOIN workspace_guest ON workspace_guest.workspace_id=workspace.id
		WHERE workspace_guest.user_email=?`,
		userID)
	if err != nil {
		return nil, fmt.Errorf("Cannot get workspaces having user ID %s as guest: %w", userID, err)
	}
	defer rows.Close()
	var workspaces []*model.Workspace
	for rows.Next() {
		workspace := model.Workspace{
			Owner: model.User{},
		}
		if err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.Owner.ID, &workspace.Owner.Email); err != nil {
			return nil, err
		}
		workspaces = append(workspaces, &workspace)
	}
	return workspaces, nil
}

// GetWorkspaceGuests returns the list of non-owner users of a workspace
func (s *Service) GetWorkspaceGuests(workspaceID string) ([]*model.User, error) {
	rows, err := s.db.Query(
		`SELECT email, email
		FROM user
		INNER JOIN workspace_guest ON workspace_guest.user_email=user.email
		WHERE workspace_guest.workspace_id=?`,
		workspaceID)
	if err != nil {
		return nil, fmt.Errorf("Cannot get users for workspace ID %s: %w", workspaceID, err)
	}
	defer rows.Close()
	var users []*model.User
	for rows.Next() {
		user := model.User{}
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
