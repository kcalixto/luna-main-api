run: clean build deploy

install-dependencies:
	npm ci

build:
	GOOS=linux GOARCH=amd64 go build -o bin/server server/handler/main.go

clean:
	rm -rf ./bin

deploy:	
	serverless create-cert --stage $(STAGE) --verbose
	serverless create_domain --stage $(STAGE) --verbose
	serverless deploy --stage $(STAGE) --verbose