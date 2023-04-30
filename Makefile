run: clean build deploy

build:
	export GO111MODULE=on
	export CGO_ENABLED=0

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/server server/handler/main.go

clean:
	rm -rf ./bin

deploy:
	serverless deploy --stage $(STAGE) --verbose

deploy-infra:
	cd infra && serverless deploy --stage $(STAGE) --verbose