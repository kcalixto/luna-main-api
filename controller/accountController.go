package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	logger "github.com/lunaorg/luna-main-api/libs/log"
	"github.com/lunaorg/luna-main-api/types/viewmodel"
)

func (c *Controller) SaveAccount(e echo.Context) (err error) {
	var data types.SaveAccount
	err = c.ParseEchoContext(e, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = c.accountSvc.Save(data)
	if err != nil {
		logger.Error(err.Error())
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusCreated, "")
}

func (c *Controller) NewIncome(e echo.Context) (err error) {
	var data types.AccountTransaction
	err = c.ParseEchoContext(e, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = c.accountSvc.NewIncome(data)
	if err != nil {
		logger.Error(err.Error())
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusCreated, "")
}

func (c *Controller) NewExpense(e echo.Context) (err error) {
	var data types.AccountTransaction
	err = c.ParseEchoContext(e, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = c.accountSvc.NewExpense(data)
	if err != nil {
		logger.Error(err.Error())
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusCreated, "")
}
