package services

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	controller "github.com/lunaorg/luna-main-api/controller"
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

func (s *Server) StartLocal(port string) {
	s.Mount()
	s.e.Start(port)
}

func (s *Server) Mount() {
	endpoint, err := controller.NewController()
	if err != nil {
		panic(err)
	}

	s.e.POST("/v1/login", endpoint.Token)
	s.e.POST("/v1/register", endpoint.RegisterUser)

	s.e.POST("/v1/new-account", endpoint.SaveAccount)
	s.e.POST("/v1/account/new-income", endpoint.NewIncome)
	s.e.POST("/v1/account/new-expense", endpoint.NewExpense)
}
