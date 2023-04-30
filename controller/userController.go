package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	logger "github.com/lunaorg/luna-main-api/libs/log"
	"github.com/lunaorg/luna-main-api/types/viewmodel"
)

func (c *Controller) RegisterUser(e echo.Context) (err error) {
	var data types.RegisterUser
	err = c.ParseEchoContext(e, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	svc, err := c.service.NewUserService()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = svc.RegisterUser(data)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	msg := fmt.Sprintf("Welcome, %s", data.Login)
	return e.JSON(http.StatusOK, msg)
}
