package auth

import (
	"crypto/sha256"
	"net/http"
	"time"

	"github.com/blacksfk/are_hub"
	"github.com/blacksfk/uf"
	"go.mozilla.org/hawk"
)

// Check whether or not the provided nonce has been used previously.
func validNonce(nonce string, t time.Time, creds *hawk.Credentials) bool {
	// TODO: ensure nonce has only been used once
	return true
}

// Authentication middleware.
type Authentication struct {
	users are_hub.UserRepo
}

// Create a new user authentication service.
func New(users are_hub.UserRepo) Authentication {
	return Authentication{users}
}

// Authenticate incoming requests with Hawk.
func (a Authentication) Request(r *http.Request) error {
	var user *are_hub.User
	ctx := r.Context()

	// authenticate the request
	hawkAuth, e := hawk.NewAuthFromRequest(r, func(creds *hawk.Credentials) error {
		var e error

		// find the user
		user, e = a.users.FindID(ctx, creds.ID)

		if e != nil {
			if are_hub.IsNoObjectsFound(e) {
				return uf.NotFound(e.Error())
			}

			return e
		}

		// set the key and hash algorithm
		creds.Key = user.Key
		creds.Hash = sha256.New

		return nil
	}, validNonce)

	if e != nil {
		if _, ok := e.(hawk.AuthError); ok {
			return uf.BadRequest(e.Error())
		}

		return e
	}

	// is the request valid?
	e = hawkAuth.Valid()

	if e != nil {
		if _, ok := e.(hawk.AuthError); ok {
			return uf.Unauthorized(e.Error())
		}

		return e
	}

	// no error returned so the user is authenticated
	*r = *(r.WithContext(user.SetAuthenticated(ctx)))

	return nil
}
