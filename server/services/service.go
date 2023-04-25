package services

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	_controller "github.com/lunaorg/luna-main-api/controller"
)

type Server struct {
	e *echo.Echo
	l *echoadapter.EchoLambda
}

func NewServer() *Server {
	e := echo.New()
	return &Server{
		e: e,
		l: echoadapter.New(e),
	}
}

func (s *Server) ProxyWithContext(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	s.Mount()
	return s.l.ProxyWithContext(ctx, req)
}

func (s *Server) Mount() {
	controller := _controller.NewController()

	s.e.POST("/v1/login", controller.TokenController)
	s.e.POST("/v1/register", controller.RegisterUser)
}
