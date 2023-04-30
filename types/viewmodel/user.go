package types

import (
	"regexp"
)

type RegisterUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (r *RegisterUser) Validate() (e []string) {
	if r.Login == "" {
		e = append(e, "empty `login`")
	}
	if match, _ := regexp.MatchString(`^[a-z0-9]+(?:(?:-|_)+[a-z0-9]+)*$`, r.Login); !match {
		e = append(e, "login must not contain spaces or special characters")
	}
	if r.Password == "" {
		e = append(e, "empty `password`")
	}
	if r.Name == "" {
		e = append(e, "empty `name`")
	}

	return e
}
