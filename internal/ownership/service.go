package ownership

import "github.com/tiramiseb/budbud-api/internal/ownership/model"

// Storage is storage for ownership
type Storage interface {
	AddWorkspace(userID, name string) (*model.Workspace, error)
	GetWorkspaceForUser(userID, id string) (*model.Workspace, error)
	GetAllWorkspacesForUser(userID string) ([]*model.Workspace, error)
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

// GetWorkspaceForUser returns all workspaces a user has access to
func (s *Service) GetWorkspaceForUser(userID, id string) (*model.Workspace, error) {
	return s.stor.GetWorkspaceForUser(userID, id)
}

// GetAllWorkspacesForUser returns all workspaces a user has access to
func (s *Service) GetAllWorkspacesForUser(userID string) ([]*model.Workspace, error) {
	return s.stor.GetAllWorkspacesForUser(userID)
}

// GetWorkspaceGuests returns all non-owner users of a workspace
func (s *Service) GetWorkspaceGuests(workspaceID string) ([]*model.User, error) {
	return s.stor.GetWorkspaceGuests(workspaceID)
}
