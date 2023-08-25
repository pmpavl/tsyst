PROJECT?=prokmak
APP?=tsyst-service
PORT?=80
PORT_APP?=7784

CONTAINER_IMAGE?=$(PROJECT)/${APP}
RELEASE?=0.0.2

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

container-build:
	@docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

container-run:
	@docker run --name ${APP} -p ${PORT}:${PORT_APP} --rm \
		-e "PORT=${PORT}" \
		--env-file .env.prod \
		$(CONTAINER_IMAGE):$(RELEASE)
