# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

WORKDIR /app

COPY ../go.mod .
COPY ../go.sum .
COPY ../app ./app
COPY ../cmd ./cmd
COPY ../configs ./configs
COPY ../pkg ./pkg

RUN go mod download

ARG SERVICE_NAME
RUN go build ./cmd/$SERVICE_NAME/main.go

ARG PORT
EXPOSE $PORT

CMD ["/app/main"]
