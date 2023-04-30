package service

import (
	data "github.com/lunaorg/luna-main-api/data/user"
	logger "github.com/lunaorg/luna-main-api/libs/log"
	"github.com/lunaorg/luna-main-api/types/viewmodel"
)

type UserService struct {
	userData *data.UserData
}

func (s *Service) NewUserService() (*UserService, error) {
	userData, err := data.NewUserData()
	if err != nil {
		return nil, err
	}

	return &UserService{
		userData: userData,
	}, nil
}

func (s *UserService) RegisterUser(user types.RegisterUser) error {
	err := s.userData.Register(user)
	if err != nil {
		logger.Error(err.Error())
	}

	return nil
}
