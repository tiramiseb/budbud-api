package gqlgen

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/rs/cors"

	"github.com/tiramiseb/budbud-api/internal/delivery/gqlgen/generated"
)

// Service provides a GraphQL server
type Service struct {
	httpServer *http.Server

	authn     authnSrv
	ownership ownershipSrv
}

// New returns a new GraphQL server
func New(port int, a authnSrv, o ownershipSrv) (*Service, error) {
	service := &Service{
		authn:     a,
		ownership: o,
	}
	config := generated.Config{Resolvers: service}
	config.Directives.Auth = authenticationDirective
	serveMux := http.NewServeMux()
	serveMux.Handle("/", handler.Playground("Bud-Bud GraphQL playground", "/query"))
	serveMux.Handle(
		"/query",
		// TODO Remove in production below this line
		cors.New(cors.Options{
			AllowOriginFunc:  func(string) bool { return true },
			AllowCredentials: true,
		}).Handler(
			// TODO Remove in production above this line
			authnMiddleware(a)(
				handler.GraphQL(generated.NewExecutableSchema(config)),
			),
		), // TODO Remove in production
	)
	service.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: serveMux,
	}
	return service, nil
}

// Start starts the GraphQL server
func (s *Service) Start() error {
	return s.httpServer.ListenAndServe()
}

type mutation struct {
	srv *Service
}

// Mutation returns the mutation resolver
func (s *Service) Mutation() generated.MutationResolver {
	return &mutation{s}
}

type query struct {
	srv *Service
}

// Query returns the query resolver
func (s *Service) Query() generated.QueryResolver {
	return &query{s}
}
