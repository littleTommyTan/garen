FROM golang:latest

WORKDIR $GOPATH/src/github.com/tommytan/tommytan-gin
COPY . $GOPATH/src/github.com/tommytan/tommytan-gin

RUN go build .

EXPOSE 80
CMD ["./garen"]
