APP?=tsyst

clean:
	@rm -f bin/${APP}

test:
	@go test -tags="testing" -v -race -cover -coverprofile=coverage.out ./...

cover: test
	@go tool cover -html=coverage.out

linter:
	@golangci-lint run

build: clean
	@go build \
		-tags go_json \
		-o bin/${APP} \
		./cmd/

gorun: build
	@bin/${APP}

docker-compose-up:
	@docker-compose up -d --build

docker-compose-down:
	@docker-compose down

docker-compose-up-mongo:
	@docker-compose up -d mongo
