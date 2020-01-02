package gqlgen

import (
	"github.com/tiramiseb/budbud-api/internal/delivery/gqlgen/generated"
)

type workspaceOwned struct {
	srv *Service
}

// WorkspaceOwned returns the owned workspace resolver
func (s *Service) WorkspaceOwned() generated.WorkspaceOwnedResolver {
	return &workspaceOwned{s}
}
