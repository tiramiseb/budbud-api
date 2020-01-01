package gqlgen

import (
	"context"
	"errors"

	"github.com/tiramiseb/budbud-api/internal/ownership/model"
)

type authnSrv interface {
	GetUserFromToken(token string) (model.User, error)
	CreateTokenFromCredentials(email, password string) (string, model.User, error)
	RemoveToken(token string) error
}

func (q *query) Me(ctx context.Context) (*model.User, error) {
	user := CurrentUser(ctx)
	return user, nil
}

func (m *mutation) Login(ctx context.Context, email string, password string) (*model.User, error) {
	token, user, err := m.srv.authn.CreateTokenFromCredentials(email, password)
	if err != nil {
		return nil, err
	}
	r := *(ResponseWriter(ctx))
	r.Header().Add("Set-Cookie", "budbud-authn="+token+"; Max-Age: 21600; Path=/query; HttpOnly")
	return &user, nil
}

func (m *mutation) Logout(ctx context.Context, none *bool) (*bool, error) {
	r := *(ResponseWriter(ctx))
	r.Header().Add("Set-Cookie", "budbud-authn=deleted; Max-Age: 1; Path=/query; HttpOnly")
	token := CurrentToken(ctx)
	if token == "" {
		response := false
		return &response, errors.New("No token defined")
	}
	err := m.srv.authn.RemoveToken(token)
	if err != nil {
		return nil, err
	}
	response := err == nil
	return &response, err
}
