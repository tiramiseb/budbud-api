package gqlgen

import (
	"github.com/tiramiseb/budbud-api/internal/delivery/gqlgen/generated"
)

type workspace struct {
	srv *Service
}

// Workspace returns the owned workspace resolver
func (s *Service) Workspace() generated.WorkspaceResolver {
	return &workspace{s}
}
