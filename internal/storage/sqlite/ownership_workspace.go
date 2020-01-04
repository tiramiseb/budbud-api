package sqlite

import (
	"fmt"
	"strconv"

	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

// AddWorkspace adds a new workspace
func (s *Service) AddWorkspace(userID, name string) (*model.Workspace, error) {
	result, err := s.db.Exec("INSERT INTO workspace (owner_email, name) VALUES (?, ?)", userID, name)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	return s.GetWorkspaceForUser(userID, strconv.FormatInt(id, 10))
}

// GetWorkspaceForUser returns the given workspace
func (s *Service) GetWorkspaceForUser(userID, id string) (*model.Workspace, error) {
	workspace := model.Workspace{
		Owner: model.User{},
	}
	err := s.db.QueryRow(
		`SELECT workspace.id, workspace.name, user.email, user.email
		FROM workspace
		INNER JOIN user ON workspace.owner_email=user.email
		LEFT JOIN workspace_guest ON workspace_guest.workspace_id=workspace.id
		WHERE (workspace.owner_email=? OR workspace_guest.user_email=?) AND workspace.id=?`,
		userID, userID, id,
	).Scan(&workspace.ID, &workspace.Name, &workspace.Owner.ID, &workspace.Owner.Email)
	return &workspace, err
}

// GetAllWorkspacesForUser returns the list of workspaces a user has access to
func (s *Service) GetAllWorkspacesForUser(userID string) ([]*model.Workspace, error) {
	rows, err := s.db.Query(
		`SELECT workspace.id, workspace.name, user.email, user.email
		FROM workspace
		INNER JOIN user ON workspace.owner_email=user.email
		LEFT JOIN workspace_guest ON workspace_guest.workspace_id=workspace.id
		WHERE workspace.owner_email=? OR workspace_guest.user_email=?`,
		userID, userID)
	if err != nil {
		return nil, fmt.Errorf("Cannot get workspaces for user ID %s: %w", userID, err)
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
