package validate

import (
	"net/http"

	"github.com/blacksfk/are_hub"
	"github.com/go-playground/validator/v10"
)

type Channel struct {
	request
}

// Create a new channel validator which exports methods matching the
// microframework.Middleware signature.
func NewChannel() Channel {
	// dependency injection? never heard of it...
	return Channel{request{validator.New()}}
}

type channelStore struct {
	Name            string `validate:"required"`
	Password        string `validate:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `validate:"required"`
}

// Validate the request body with rules defined above. If successful,
// create an are_hub.Channel and attach it to the request's context.
func (c Channel) Store(r *http.Request) error {
	temp := channelStore{}
	e := c.bodyStruct(r, &temp)

	if e != nil {
		return e
	}

	// create a channel (domain type) out of the validation object
	channel := are_hub.NewChannel(temp.Name, temp.Password)

	// insert the create channel into r's context
	*r = *r.WithContext(channel.ToCtx(r.Context()))

	return nil
}

type verifyPassword struct {
	Password string `validate:"required"`
}

// Validate the request body with the rules defined in verifyPassword.
// If successful, create a channel and embed it into the request's context.
func (c Channel) Verify(r *http.Request) error {
	temp := verifyPassword{}
	e := c.bodyStruct(r, &temp)

	if e != nil {
		return e
	}

	// create a channel out of the validation object
	channel := are_hub.NewChannel("", temp.Password)

	// insert the channel into r's context
	*r = *r.WithContext(channel.ToCtx(r.Context()))

	return nil
}
