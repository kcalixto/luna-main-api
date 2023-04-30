package types

import (
	"errors"
	"regexp"
)

type RegisterUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (r *RegisterUser) Validate() (e []error) {
	if r.Login == "" {
		e = append(e, errors.New("empty Login"))
	}
	if match, _ := regexp.MatchString(`^[a-z0-9]+(?:(?:-|_)+[a-z0-9]+)*$`, r.Login); !match {
		e = append(e, errors.New("login must not contain spaces or special characters"))
	}
	if r.Password == "" {
		e = append(e, errors.New("empty Password"))
	}

	return e
}
