FROM golang:latest

ENV GIN_MODE release
WORKDIR $GOPATH/src/github.com/tommytan/garen
COPY . $GOPATH/src/github.com/tommytan/garen

RUN go build .

EXPOSE 80
CMD ["./garen"]
