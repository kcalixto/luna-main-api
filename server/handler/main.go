package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lunaorg/luna-main-api/server/services"
)

var svc *services.Server

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("raw request: %+v \n", req)
	}

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

		item := map[string]interface{}{
			"login":    "teste",
			"password": "local",
		}

		bytes, err := json.Marshal(item)
		if err != nil {
			panic(err)
		}

		req := events.APIGatewayProxyRequest{
			HTTPMethod: "POST",
			Path:       "/login",
			Resource:   "/login",
			Body:       string(bytes),
		}

		fmt.Println(Handler(context.Background(), req))
	} else {
		lambda.Start(Handler)
	}
}
