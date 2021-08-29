package http

import (
	"net/http"

	"github.com/blacksfk/are_hub"
	"github.com/blacksfk/are_hub/hash"
	"github.com/blacksfk/uf"
)

type User struct {
	users are_hub.UserRepo
}

// Create a new user controller.
func NewUser(users are_hub.UserRepo) User {
	return User{users}
}

// Get all users without their keys and passwords.
func (u User) Index(w http.ResponseWriter, r *http.Request) error {
	users, e := u.users.All(r.Context())

	if e != nil {
		return e
	}

	return uf.SendJSON(w, users)
}

// Create a new user.
func (u User) Store(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	user, e := are_hub.UserFromCtx(ctx)

	if e != nil {
		return e
	}

	// hash the new user's password
	hashed, e := hash.Password(user.Password.String())

	if e != nil {
		return e
	}

	user.Password.Set(hashed)
	e = u.users.Insert(ctx, user)

	if e != nil {
		return e
	}

	return uf.SendJSON(w, user)
}

// Find a specific channel by its ID.
func (u User) Show(w http.ResponseWriter, r *http.Request) error {
	user, e := u.users.FindID(r.Context(), uf.GetParam(r, "id"))

	if e != nil {
		return e
	}

	return uf.SendJSON(w, user)
}

// Update a specific user by its ID.
func (u User) Update(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(501)

	return nil
}

// Delete a specific user by its ID.
func (u User) Delete(w http.ResponseWriter, r *http.Request) error {
	user, e := u.users.DeleteID(r.Context(), uf.GetParam(r, "id"))

	if e != nil {
		return e
	}

	return uf.SendJSON(w, user)
}

// Login with a username and password.
func (u User) Login(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	login, e := are_hub.LoginFromCtx(ctx)

	if e != nil {
		return e
	}

	// find the user by their username
	user, e := u.users.FindName(ctx, login.Name)

	if e != nil {
		if are_hub.IsNoObjectsFound(e) {
			return uf.Unauthorized("Incorrect username or password")
		}

		return e
	}

	// compare the known and provided passwords
	match, e := hash.CmpPassword(user.Password.String(), login.Password)

	if e != nil {
		return e
	}

	if !match {
		return uf.Unauthorized("Incorrect username or password")
	}

	// passwords match; generate a new key
	user.Key, e = hash.GenerateKey()

	if e != nil {
		return e
	}

	// update the user in the db with the new key
	e = u.users.UpdateID(ctx, user.ID, user)

	if e != nil {
		return e
	}

	// login success; send the user object with their ID and key
	return uf.SendJSON(w, user)
}
