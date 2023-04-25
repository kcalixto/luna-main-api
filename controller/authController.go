package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lunaorg/luna-main-api/types"
)

func (c *Controller) TokenController(e echo.Context) error {
	var data types.AuthTokenInput
	err := c.ParseEchoContext(e, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_ = c.service.NewAuthService(data)

	return e.JSON(http.StatusOK, fmt.Sprintf("hello %s", data.Login))
}
