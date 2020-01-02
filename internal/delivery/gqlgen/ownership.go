package gqlgen

import (
	"context"

	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

type ownershipSrv interface {
	AddWorkspace(userID, name string) (*model.Workspace, error)
	GetWorkspacesOwned(userID string) ([]*model.Workspace, error)
	GetWorkspacesGuest(userID string) ([]*model.Workspace, error)
	GetWorkspaceGuests(workspaceID string) ([]*model.User, error)
}

func (m *mutation) AddWorkspace(ctx context.Context, name string) (*model.Workspace, error) {
	user := CurrentUser(ctx)
	return m.srv.ownership.AddWorkspace(user.ID, name)
}

func (q *query) WorkspacesOwned(ctx context.Context) ([]*model.Workspace, error) {
	user := CurrentUser(ctx)
	return q.srv.ownership.GetWorkspacesOwned(user.ID)
}

func (q *query) WorkspacesGuest(ctx context.Context) ([]*model.Workspace, error) {
	user := CurrentUser(ctx)
	return q.srv.ownership.GetWorkspacesGuest(user.ID)
}

func (w *workspaceOwned) Guests(ctx context.Context, workspace *model.Workspace) ([]*model.User, error) {
	return w.srv.ownership.GetWorkspaceGuests(workspace.ID)
}
