package gqlgen

import (
	"context"

	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

type ownershipSrv interface {
	AddWorkspace(userID, name string) (*model.Workspace, error)
	GetWorkspaceForUser(userID, id string) (*model.Workspace, error)
	GetAllWorkspacesForUser(userID string) ([]*model.Workspace, error)
	GetWorkspaceGuests(workspaceID string) ([]*model.User, error)
}

func (m *mutation) AddWorkspace(ctx context.Context, name string) (*model.Workspace, error) {
	user := CurrentUser(ctx)
	return m.srv.ownership.AddWorkspace(user.ID, name)
}

func (q *query) Workspace(ctx context.Context, id string) (*model.Workspace, error) {
	user := CurrentUser(ctx)
	return q.srv.ownership.GetWorkspaceForUser(user.ID, id)
}

func (q *query) Workspaces(ctx context.Context) ([]*model.Workspace, error) {
	user := CurrentUser(ctx)
	return q.srv.ownership.GetAllWorkspacesForUser(user.ID)
}

func (w *workspace) Guests(ctx context.Context, workspace *model.Workspace) ([]*model.User, error) {
	return w.srv.ownership.GetWorkspaceGuests(workspace.ID)
}
