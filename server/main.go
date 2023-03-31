package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Server struct{}

func (s *Server) Mount(e *echo.Echo) {
	log.Printf("e: %+v", e)
	
	e.GET("/hello", s.GetHello)
	e.POST("/hello", s.PostHello)
	e.GET("/greeting", s.GetGreeting)
	e.POST("/address", s.PostAddress)
}

func (s *Server) GetHello(c echo.Context) error {
	log.Print(c.Request())

	return c.JSON(http.StatusOK, "Hello World")
}

func (s *Server) PostHello(c echo.Context) error {
	log.Print("MASUK HELLO NIIIIH")
	log.Print(c.Request())
	log.Print(c.ParamValues())
	log.Print(c.FormParams())
	log.Print(c.Request().Body)

	type Test struct {
		Name string `json:"name"`
	}

	var req Test

	requestBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(requestBytes, &req)
	if err != nil {
		panic(err)
	}

	hello := "Hello, " + req.Name + "!"

	return c.JSON(http.StatusOK, hello)
}

func (s *Server) GetGreeting(c echo.Context) error {
	return c.JSON(http.StatusOK, "Welcome to test")
}

func (s *Server) PostAddress(c echo.Context) error {
	address := c.FormValue("address")
	txt := "Your address is " + address

	return c.JSON(http.StatusOK, txt)
}
