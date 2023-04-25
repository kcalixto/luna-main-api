package service

import (
	"fmt"
	"github.com/lunaorg/luna-main-api/types"
)

type UserService struct {
	data types.RegisterUser
}

func (s *Service) NewUserService(data types.RegisterUser) *UserService {
	return &UserService{
		data: data,
	}
}

func (s *UserService) RegisterUser() error {
	fmt.Printf("\nregistered. %+v", s.data)

	return nil
}
