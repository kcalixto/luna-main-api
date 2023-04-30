run: clean build deploy

install-dependencies:
	npm ci

build:
	export GO111MODULE=on
	export CGO_ENABLED=1

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/server server/handler/main.go

clean:
	rm -rf ./bin

deploy:
	echo "debug - show files"
	pwd
	ls -R
	echo "*----------------*"
	serverless deploy --stage $(STAGE) --verbose

deploy-infra:
	cd infra && serverless deploy --stage $(STAGE) --verbose