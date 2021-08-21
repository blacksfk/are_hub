package are_hub

import (
	"context"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Password password
	Key      string `json:"key"`
	Common   `bson:",inline"`
}

// Create a new user.
func NewUser(name, pw string) *User {
	return &User{Name: name, Password: NewPassword(pw)}
}

// Extract either the authenticated user or user object embedded in the context.
func userFromCtx(ctx context.Context, key ctxKey) (*User, error) {
	v := ctx.Value(key)
	u, ok := v.(*User)

	if !ok {
		return nil, fmt.Errorf("Could not assert %v as a *User", v)
	}

	return u, nil
}

// Get a user from a context.
func UserFromCtx(ctx context.Context) (*User, error) {
	return userFromCtx(ctx, keyUser)
}

// Get the authenticated user from a context.
func GetAuthenticatedUser(ctx context.Context) (*User, error) {
	return userFromCtx(ctx, keyAuth)
}

// Insert a user into a context.
func (u *User) ToCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, keyUser, u)
}

// Set the user as the authenticated user.
func (u *User) SetAuthenticated(ctx context.Context) context.Context {
	return context.WithValue(ctx, keyAuth, u)
}

type UserRepo interface {
	// Get all users.
	All(context.Context) ([]User, error)

	// Create a new user.
	Insert(context.Context, Archetype) error

	// Find a user by its ID.
	FindID(context.Context, string) (*User, error)

	// Find and update a user by its ID.
	UpdateID(context.Context, string, Archetype) error

	// Find and delete a user by its ID.
	DeleteID(context.Context, string) (*User, error)

	// Get a count of users.
	Count(context.Context) (int64, error)
}
