package ownership

import "github.com/tiramiseb/budbud-api/internal/ownership/model"

// Storage is storage for ownership
type Storage interface {
	GetWorkspacesForUserID(userID string) ([]*model.Workspace, error)
}

// Service is an ownership service
type Service struct {
	stor Storage
}

// New returns a new ownership service
func New(s Storage) (*Service, error) {
	return &Service{s}, nil
}

// GetWorkspaces returns all workspaces owned by a user
func (s *Service) GetWorkspaces(userID string) ([]*model.Workspace, error) {
	return s.stor.GetWorkspacesForUserID(userID)
}
