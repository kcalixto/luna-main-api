package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo"
	server "github.com/lunaorg/luna-main-api/server"
)

func main() {
	e := echo.New()

	newServer := server.Server{}
	newServer.Mount(e)

	isLambda := os.Getenv("LAMBDA")

	if isLambda == "TRUE" {
		lambdaAdapter := &server.LambdaAdapter{Echo: e}
		lambda.Start(lambdaAdapter.Handler)
	} else {
		e.Logger.Fatal(e.Start(":1234"))
	}
}
