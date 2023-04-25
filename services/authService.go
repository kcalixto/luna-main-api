package service

import (
	"fmt"
	"github.com/lunaorg/luna-main-api/types"
)

type AuthService struct {
	login string
}

func (s *Service) NewAuthService(auth types.AuthTokenInput) *AuthService {
	return &AuthService{
		login: auth.Login,
	}
}

func (s *AuthService) PrintName() {
	fmt.Println(s.login)
}
