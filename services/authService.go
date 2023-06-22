package service

import (
	"fmt"
	"github.com/lunaorg/luna-main-api/types"
)

type AuthService struct{}

func NewAuthService() (*AuthService, error) {
	return &AuthService{}, nil
}

func (s *AuthService) PrintName(auth types.AuthTokenInput) {
	fmt.Println(auth.Login)
}
