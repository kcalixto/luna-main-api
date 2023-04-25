package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lunaorg/luna-main-api/types"
)

func (c *Controller) RegisterUser(e echo.Context) (err error) {
	var data types.RegisterUser
	err = c.ParseEchoContext(e, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	svc := c.service.NewUserService(data)
	err = svc.RegisterUser()
	if err != nil {
		fmt.Println(err)
		return err
	}

	msg := fmt.Sprintf("Welcome, %s", data.Login)
	return e.JSON(http.StatusOK, msg)
}

// func GetGreetingController(e echo.Context) error {
// 	return e.JSON(http.StatusOK, "Welcome to test")
// }

// func PostAddressController(e echo.Context) error {
// 	address := e.FormValue("address")
// 	txt := "Your address is " + address

// 	return e.JSON(http.StatusOK, txt)
// }
