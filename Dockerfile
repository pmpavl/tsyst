FROM golang:latest as build
COPY go.mod go.sum /src/
WORKDIR /src
RUN go mod download
COPY . /src/
RUN	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/srv ./cmd

FROM alpine:latest as production
COPY --from=build /bin/srv /app/srv
RUN apk --update add ca-certificates htop
ENV PORT 7784
EXPOSE $PORT
ENTRYPOINT /app/srv
