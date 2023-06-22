package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lunaorg/luna-main-api/server/services"
)

var svc *services.Server

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return svc.ProxyWithContext(ctx, req)
}

func initDependencies() (err error) {
	svc = services.NewServer()

	return nil
}

func main() {
	err := initDependencies()
	if err != nil {
		panic(err)
	}

	if os.Getenv("NODE_ENV") == "local" {
		fmt.Println("running local")
		svc.StartLocal(":8080")
	} else {
		lambda.Start(Handler)
	}
}
