package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo"

	// "github.com/labstack/echo/middleware"
	_controller "github.com/lunaorg/luna-main-api/controller"
	"github.com/lunaorg/luna-main-api/server/adapter"
)

func Mount(e *echo.Echo) {
	controller := _controller.NewController()

	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey:  []byte(os.Getenv("JWT_SECRET")),
	// 	TokenLookup: "Authorization",
	// 	Skipper: func(c echo.Context) bool {
	// 		return c.Request().URL.Path == "/login"
	// 	},
	// }))

	e.POST("/login", controller.TokenController)

	//* user
	e.POST("/register", controller.RegisterUser)
}

func main() {
	e := echo.New()

	Mount(e)

	if os.Getenv("NODE_ENV") == "local" {
		fmt.Println("running_local")
		e.Logger.Fatal(e.Start(":3000"))
	} else {
		lambdaAdapter := &adapter.LambdaAdapter{Echo: e}
		lambda.Start(lambdaAdapter.Handler)
	}
}
