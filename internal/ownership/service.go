package ownership

import "github.com/tiramiseb/budbud-api/internal/ownership/model"

// Storage is storage for ownership
type Storage interface {
	AddWorkspace(userID, name string) (*model.Workspace, error)
	GetWorkspacesOwned(userID string) ([]*model.Workspace, error)
	GetWorkspacesGuest(userID string) ([]*model.Workspace, error)
	GetWorkspaceGuests(workspaceID string) ([]*model.User, error)
}

// Service is an ownership service
type Service struct {
	stor Storage
}

// New returns a new ownership service
func New(s Storage) (*Service, error) {
	return &Service{s}, nil
}

// AddWorkspace adds a workspace
func (s *Service) AddWorkspace(userID, name string) (*model.Workspace, error) {
	return s.stor.AddWorkspace(userID, name)
}

// GetWorkspacesOwned returns all workspaces owned by a user
func (s *Service) GetWorkspacesOwned(userID string) ([]*model.Workspace, error) {
	return s.stor.GetWorkspacesOwned(userID)
}

// GetWorkspacesGuest returns all workspaces a user has access to as a guest
func (s *Service) GetWorkspacesGuest(userID string) ([]*model.Workspace, error) {
	return s.stor.GetWorkspacesGuest(userID)
}

// GetWorkspaceGuests returns all non-owner users of a workspace
func (s *Service) GetWorkspaceGuests(workspaceID string) ([]*model.User, error) {
	return s.stor.GetWorkspaceGuests(workspaceID)
}
