run: clean build deploy

build:
	go build -ldflags="-s -w" -o bin/server server/handler/main.go

clean:
	rm -rf ./bin

deploy:
	serverless deploy --stage $(STAGE) --verbose

deploy-infra:
	cd infra && serverless deploy --stage $(STAGE) --verbose