package are_hub

import (
	"context"
	"fmt"
)

// A login represents a request to login with a username and password.
type Login struct {
	Name     string
	Password string
}

// Create a new login.
func NewLogin(name, password string) *Login {
	return &Login{Name: name, Password: password}
}

// Get a login from a context.
func LoginFromCtx(ctx context.Context) (*Login, error) {
	v := ctx.Value(keyLogin)
	l, ok := v.(*Login)

	if !ok {
		return nil, fmt.Errorf("Could not assert %v as *Login", v)
	}

	return l, nil
}

// Embed the login in the context.
func (l *Login) ToCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, keyLogin, l)
}
