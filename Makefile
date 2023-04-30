run: clean build deploy

install-dependencies:
	npm ci

build:
	GOOS=linux GOARCH=amd64 go build -o bin/server server/handler/main.go

clean:
	rm -rf ./bin

deploy:	
	serverless deploy --stage $(STAGE) --verbose

deploy-infra:	
	cd infra && serverless deploy --stage $(STAGE) --verbose