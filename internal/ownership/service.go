package ownership

import (
	"errors"

	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

// Storage is storage for ownership
type Storage interface {
	AddWorkspace(userID, name string) (*model.Workspace, error)
	GetWorkspaceForUserByID(userID, id string) (*model.Workspace, error)
	GetWorkspaceForUserByOwnerIDAndName(userID, ownerID, name string) (*model.Workspace, error)
	GetWorkspaceForUserByOwnerEmailAndName(userID, ownerEmail, name string) (*model.Workspace, error)
	GetAllWorkspacesForUser(userID string) ([]*model.Workspace, error)
	GetWorkspaceGuests(workspaceID string) ([]*model.User, error)
	GetWorkspaceCategories(workspaceID string) ([]*model.SuperCategory, error)
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
	if name == "" {
		return nil, errors.New("Name cannot be empty")
	}
	return s.stor.AddWorkspace(userID, name)
}

// GetWorkspaceForUser returns all workspaces a user has access to
//
// One of the following must be provided (resolution order is the order of the list):
// * id to return a workspace identified by its ID
// * ownerID and name to return a corresponding workspace, accessible to the giben userID
// * ownerEmail and name to return a corresponding workspace, accessible to the giben userID
// * name to return a workspace with that name owned by the given userID
func (s *Service) GetWorkspaceForUser(userID, id, ownerID, ownerEmail, name string) (*model.Workspace, error) {
	if id != "" {
		return s.stor.GetWorkspaceForUserByID(userID, id)
	}
	if name == "" {
		return nil, errors.New("ID and name cannot be both empty")
	}
	if ownerEmail == "" {
		return s.stor.GetWorkspaceForUserByOwnerIDAndName(userID, userID, name)
	}
	return s.stor.GetWorkspaceForUserByOwnerEmailAndName(userID, ownerEmail, name)
}

// GetAllWorkspacesForUser returns all workspaces a user has access to
func (s *Service) GetAllWorkspacesForUser(userID string) ([]*model.Workspace, error) {
	return s.stor.GetAllWorkspacesForUser(userID)
}

// GetWorkspaceGuests returns all non-owner users of a workspace
func (s *Service) GetWorkspaceGuests(workspaceID string) ([]*model.User, error) {
	return s.stor.GetWorkspaceGuests(workspaceID)
}

// GetWorkspaceCategories returns all categories of a workspace
func (s *Service) GetWorkspaceCategories(workspaceID string) ([]*model.SuperCategory, error) {
	return s.stor.GetWorkspaceCategories(workspaceID)
}
