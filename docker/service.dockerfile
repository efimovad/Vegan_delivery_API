FROM golang:1.14.0

WORKDIR /go/src/api
COPY . .

RUN go build -o api.app internal/main.go

ENTRYPOINT ["./api.app"]
