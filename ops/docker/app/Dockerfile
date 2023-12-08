FROM golang:1.21-alpine

WORKDIR /go/src
COPY ./ ./

WORKDIR /go/src/app

RUN go mod download

RUN go install github.com/cosmtrek/air@latest
RUN go install go.uber.org/mock/mockgen@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

CMD ["air", "-c", ".air.toml"]
