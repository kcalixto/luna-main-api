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

	err = c.userSvc.RegisterUser(data)
	if err != nil {
		logger.Error(err.Error())

		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusCreated, "")
}
