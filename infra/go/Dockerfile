FROM golang:1.23

WORKDIR /go-fargate-demo/src

COPY ./src/go.mod ./src/go.sum ./

RUN go mod download && go mod verify

COPY ./src .

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]