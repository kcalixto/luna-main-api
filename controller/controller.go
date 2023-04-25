package controller

import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/lunaorg/luna-main-api/services"
)

type Controller struct {
	service *service.Service
}

func NewController() *Controller {
	return &Controller{}
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