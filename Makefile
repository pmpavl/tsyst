APP?=tsyst

clean:
	@rm -f bin/${APP}

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
