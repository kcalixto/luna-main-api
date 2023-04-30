fast-deploy: clean build deploy-function
slow-deploy: clean build deploy-everything

build:
	go build -ldflags="-s -w" -o bin/server server/handler/main.go

clean:
	rm -rf ./bin

deploy-everything:
	serverless deploy --stage $(STAGE) --verbose

deploy-function:
	serverless deploy function -f $(MAIN_API_NAME) --stage $(STAGE) --verbose

deploy-infra:
	cd infra && serverless deploy --stage $(STAGE) --verbose