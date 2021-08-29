package validate

import (
	"net/http"

	"github.com/blacksfk/are_hub"
	"github.com/go-playground/validator/v10"
)

type User struct {
	request
}

func NewUser() User {
	return User{request{validator.New()}}
}

type login struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}

func (u User) Login(r *http.Request) error {
	temp := login{}
	e := u.bodyStruct(r, &temp)

	if e != nil {
		return e
	}

	// create a login out of the validation object
	l := are_hub.NewLogin(temp.Name, temp.Password)

	// embed the login into the request's context
	*r = *(r.WithContext(l.ToCtx(r.Context())))

	return nil
}
