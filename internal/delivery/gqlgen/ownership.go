package gqlgen

import (
	"context"

	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

type ownershipSrv interface {
	GetWorkspaces(userID string) ([]*model.Workspace, error)
}

func (q *query) Workspaces(ctx context.Context) ([]*model.Workspace, error) {
	user := CurrentUser(ctx)
	return q.srv.ownership.GetWorkspaces(user.ID)
}
