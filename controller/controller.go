package controller

import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
	logger "github.com/lunaorg/luna-main-api/libs/log"
	services "github.com/lunaorg/luna-main-api/services"
)

type Controller struct {
	accountSvc *services.AccountService
	userSvc    *services.UserService
	authSvc    *services.AuthService
}

func NewController() (*Controller, error) {
	authSvc, err := services.NewAuthService()
	if err != nil {
		logger.Error(err.Error())
		return &Controller{}, err
	}
	userSvc, err := services.NewUserService()
	if err != nil {
		logger.Error(err.Error())
		return &Controller{}, err
	}
	accountSvc, err := services.NewAccountService()
	if err != nil {
		logger.Error(err.Error())
		return &Controller{}, err
	}

	return &Controller{
		authSvc:    authSvc,
		userSvc:    userSvc,
		accountSvc: accountSvc,
	}, nil
}

func (c *Controller) ParseEchoContext(e echo.Context, item interface{}) error {
	requestBytes, err := io.ReadAll(e.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(requestBytes, &item)
	if err != nil {
		return err
	}

	return nil
}
